package base

import (
	"context"
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/qri/base/dsfs"
	"github.com/qri-io/qri/dsref"
	reporef "github.com/qri-io/qri/repo/ref"
)

func TestRecall(t *testing.T) {
	ctx := context.Background()
	r := newTestRepo(t)
	ref := addNowTransformDataset(t, r)

	_, err := Recall(ctx, r, "", ref)
	if err != nil {
		t.Error(err)
	}

	_, err = Recall(ctx, r, "tf", ref)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadRevisions(t *testing.T) {
	ctx := context.Background()
	r := newTestRepo(t)
	ref := addCitiesDataset(t, r)

	cities, err := dsfs.LoadDataset(ctx, r.Store(), ref.Path)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		ref  reporef.DatasetRef
		revs string
		ds   *dataset.Dataset
		err  string
	}{
		// TODO - both of these are failing, let's make 'em work:
		// "ds" Qri value is mismatching
		// {ref, "ds", cities, ""},
		// Logic on what to do in "body" is a little confusing atm, do we set BodyPath
		// and move on?
		// {ref, "bd", cities, ""},

		{ref, "tf", &dataset.Dataset{Transform: cities.Transform}, ""},
		{ref, "md,vz,tf,st", &dataset.Dataset{Transform: cities.Transform, Meta: cities.Meta, Structure: cities.Structure}, ""},
	}

	for i, c := range cases {
		revs, err := dsref.ParseRevs(c.revs)
		if err != nil {
			t.Errorf("case %d error parsing revs: %s", i, err)
			continue
		}

		got, err := LoadRevs(ctx, r, c.ref, revs)
		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}

		if err := dataset.CompareDatasets(c.ds, got); err != nil {
			t.Errorf("case %d result mismatch: %s", i, err)
		}
	}
}
