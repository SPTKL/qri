package lib

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qri-io/qri/config"
	"github.com/qri-io/qri/p2p"
	"github.com/qri-io/qri/repo"
	testrepo "github.com/qri-io/qri/repo/test"
)

func TestHistoryRequestsLog(t *testing.T) {
	mr, refs, err := testrepo.NewTestRepoWithHistory()
	if err != nil {
		t.Fatalf("error allocating test repo: %s", err.Error())
		return
	}

	node, err := p2p.NewQriNode(mr, config.DefaultP2PForTesting())
	if err != nil {
		t.Fatal(err.Error())
	}

	firstRef := refs[0].String()

	cases := []struct {
		description string
		p           *LogParams
		refs        []repo.DatasetRef
		err         string
	}{
		{"log list - empty",
			&LogParams{}, nil, "repo: empty dataset reference"},
		{"log list - bad path",
			&LogParams{Ref: "/badpath"}, nil, "repo: not found"},
		{"log list - default",
			&LogParams{Ref: firstRef}, refs, ""},
		{"log list - offset 0 limit 3",
			&LogParams{Ref: firstRef, ListParams: ListParams{Offset: 0, Limit: 3}}, refs[:3], ""},
		{"log list - offset 3 limit 3",
			&LogParams{Ref: firstRef, ListParams: ListParams{Offset: 3, Limit: 3}}, refs[3:], ""},
		{"log list - offset 6 limit 3",
			&LogParams{Ref: firstRef, ListParams: ListParams{Offset: 6, Limit: 3}}, []repo.DatasetRef{}, ""},
	}

	req := NewLogRequests(node, nil)
	for _, c := range cases {
		got := []repo.DatasetRef{}
		err := req.Log(c.p, &got)

		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case '%s' error mismatch: expected: %s, got: %s", c.description, c.err, err)
			continue
		}

		if len(c.refs) != len(got) {
			t.Errorf("case '%s' expected log length error. expected: %d, got: %d", c.description, len(c.refs), len(got))
			continue
		}

		for j, expect := range c.refs {
			if err := repo.CompareDatasetRef(expect, got[j]); err != nil {
				t.Errorf("case '%s' expected dataset error. index %d mismatch: %s", c.description, j, err.Error())
				continue
			}
		}
	}
}

func TestHistoryRequestsLogEntries(t *testing.T) {
	mr, refs, err := testrepo.NewTestRepoWithHistory()
	if err != nil {
		t.Fatalf("error allocating test repo: %s", err.Error())
		return
	}

	node, err := p2p.NewQriNode(mr, config.DefaultP2PForTesting())
	if err != nil {
		t.Fatal(err.Error())
	}

	firstRef := refs[0].String()
	req := NewLogRequests(node, nil)

	if err = req.Logbook(&RefListParams{}, nil); err == nil {
		t.Errorf("expected empty reference param to error")
	}

	res := []LogEntry{}
	if err = req.Logbook(&RefListParams{Ref: firstRef, Limit: 30}, &res); err != nil {
		t.Fatal(err)
	}

	result := make([]string, len(res))
	for i := range res {
		// set response times to zero for consistent results
		res[i].Timestamp = time.Time{}
		result[i] = res[i].String()
	}

	expect := []string{
		`12:00AM	peer	init	`,
		`12:00AM	peer	save	initial commit`,
		`12:00AM	peer	save	initial commit`,
		`12:00AM	peer	save	initial commit`,
		`12:00AM	peer	save	initial commit`,
		`12:00AM	peer	save	initial commit`,
	}

	if diff := cmp.Diff(expect, result); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
}
