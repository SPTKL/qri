package fsrepo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/qri-io/qfs/cafs"
	"github.com/qri-io/qri/repo"
)

// Refstore is a file-based implementation of the repo.Refstore
// interface. It stores names in a json file
type Refstore struct {
	basepath
	file File
	// filestore for checking dataset integrity
	store cafs.Filestore
}

// PutRef adds a reference to the store
func (rs Refstore) PutRef(r repo.DatasetRef) (err error) {
	var refs repo.Refs

	// remove dataset reference, refstores only store reference details
	r.Dataset = nil

	if r.ProfileID == "" {
		return repo.ErrPeerIDRequired
	} else if r.Name == "" {
		return repo.ErrNameRequired
	} else if r.Path == "" && r.FSIPath == "" {
		return repo.ErrPathRequired
	} else if r.Peername == "" {
		return repo.ErrPeernameRequired
	}

	if refs, err = rs.refs(); err != nil {
		return err
	}

	matched := false
	for i, ref := range refs {
		if ref.Match(r) {
			matched = true
			refs[i] = r
		}
	}

	if !matched {
		refs = append(refs, r)
	}

	return rs.save(refs)
}

// GetRef completes a partially-known reference
func (rs Refstore) GetRef(get repo.DatasetRef) (repo.DatasetRef, error) {
	refs, err := rs.refs()
	if err != nil {
		return repo.DatasetRef{}, err
	}
	for _, ref := range refs {
		if ref.Match(get) {
			return ref, nil
		}
	}
	return repo.DatasetRef{}, repo.ErrNotFound
}

// DeleteRef removes a name from the store
func (rs Refstore) DeleteRef(del repo.DatasetRef) error {
	refs, err := rs.refs()
	if err != nil {
		return err
	}

	for i, ref := range refs {
		if ref.Match(del) {
			refs = append(refs[:i], refs[i+1:]...)
			break
		}
	}

	return rs.save(refs)
}

// References gives a set of dataset references from the store
func (rs Refstore) References(offset, limit int) ([]repo.DatasetRef, error) {
	refs, err := rs.refs()
	if err != nil {
		return nil, err
	}
	res := make(repo.Refs, limit)
	for i, ref := range refs {
		if i < offset {
			continue
		}
		if i-offset == limit {
			return []repo.DatasetRef(res), nil
		}
		res[i-offset] = ref
	}
	return res[:len(refs)-offset], nil
}

// RefCount returns the size of the Refstore
func (rs Refstore) RefCount() (int, error) {
	// TODO (b5) - there's no need to unmarshal here
	// could just read the length of the flatbuffer ref vector
	refs, err := rs.refs()
	if err != nil {
		log.Debug(err.Error())
		return 0, err
	}
	return len(refs), nil
}

func (rs *Refstore) refs() (repo.Refs, error) {
	data, err := ioutil.ReadFile(rs.filepath(rs.file))
	if err != nil {
		if os.IsNotExist(err) {
			// empty is ok
			return repo.Refs{}, nil
		}
		log.Debug(err.Error())
		return nil, fmt.Errorf("error loading references: %s", err.Error())
	}

	return repo.UnmarshalRefsFlatbuffer(data)
}

func (rs *Refstore) jsonRefs() (repo.Refs, error) {
	data, err := ioutil.ReadFile(rs.filepath(rs.file))
	if err != nil {
		if os.IsNotExist(err) {
			// empty is ok
			return repo.Refs{}, nil
		}
		log.Debug(err.Error())
		return nil, fmt.Errorf("error loading names: %s", err.Error())
	}

	refs := repo.Refs{}
	err = json.Unmarshal(data, &refs)
	return refs, err
}

func (rs *Refstore) save(refs repo.Refs) error {
	sort.Sort(refs)
	path := rs.basepath.filepath(rs.file)
	return ioutil.WriteFile(path, refs.FlatbufferBytes(), os.ModePerm)
}
