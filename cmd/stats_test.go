package cmd

import (
	"testing"

	"github.com/qri-io/ioes"
)

func TestStatsComplete(t *testing.T) {
	streams, in, out, errs := ioes.NewTestIOStreams()
	setNoColor(true)

	f, err := NewTestFactory()
	if err != nil {
		t.Errorf("error creating new test factory: %s", err)
		return
	}

	badCases := []struct {
		description string
		args        []string
		err         string
	}{
		{"no args", []string{}, "need a dataset reference, eg: me/dataset_name"},
	}

	for i, c := range badCases {
		opt := &StatsOptions{
			IOStreams: streams,
		}

		err := opt.Complete(f, c.args)

		if c.err != err.Error() {
			t.Errorf("%d. case %s, error mismatch. Expected: '%s', Got: '%s'", i, c.description, err, err.Error())
			ioReset(in, out, errs)
			continue
		}

		ioReset(in, out, errs)
	}

	goodCases := []struct {
		description string
		args        []string
		expectedRef string
	}{
		{"given one ref", []string{"me/ref"}, "me/ref"},
		{"given multiple refs", []string{"me/ref", "me/ref_foo"}, "me/ref"},
	}

	for i, c := range goodCases {
		opt := &StatsOptions{
			IOStreams: streams,
		}

		if err := opt.Complete(f, c.args); err != nil {
			t.Errorf("%d. case %s, unexpected error: '%s'", i, c.description, err.Error())
		}

		if c.expectedRef != opt.Ref {
			t.Errorf("%d. case %s, incorrect ref. Expected: '%s', Got: '%s'", i, c.description, c.expectedRef, opt.Ref)
			ioReset(in, out, errs)
			continue
		}

		if opt.DatasetRequests == nil {
			t.Errorf("%d. case %s, opt.DatasetRequests not set.", i, c.description)
			ioReset(in, out, errs)
			continue
		}
		ioReset(in, out, errs)
	}
}

func TestStatsRun(t *testing.T) {
	streams, in, out, errs := ioes.NewTestIOStreams()
	setNoColor(true)
	setNoPrompt(true)

	f, err := NewTestFactory()
	if err != nil {
		t.Fatalf("error creating new test factory: %s", err)
	}

	dsr, err := f.DatasetRequests()
	if err != nil {
		t.Fatalf("error creating dataset request: %s", err)
	}
	badCases := []struct {
		description string
		ref         string
		err         string
	}{
		{"empty ref", "", "repo: empty dataset reference"},
		{"dataset does not exist", "me/dataset_does_not_exist", "repo: not found"},
	}

	for i, c := range badCases {
		ioReset(in, out, errs)

		opt := &StatsOptions{
			IOStreams:       streams,
			Ref:             c.ref,
			DatasetRequests: dsr,
		}

		if err = opt.Run(); err.Error() != c.err {
			t.Errorf("%d. case '%s', error mismatch, expected: %s, got: %s ", i, c.description, c.err, err)
			continue
		}
	}

	goodCases := []struct {
		description string
		ref         string
		// TODO (ramfox): this is cheating, we should check that the out stream is
		// getting the proper output. This will be more important when we have
		// nicely formatted stats. The veracity of the stats are being checked lower
		// in the stack
	}{
		{"csv", "me/cities"},
		{"json", "me/sitemap"},
	}

	for i, c := range goodCases {
		ioReset(in, out, errs)

		opt := &StatsOptions{
			IOStreams:       streams,
			Ref:             c.ref,
			DatasetRequests: dsr,
		}

		if err = opt.Run(); err != nil {
			t.Errorf("%d. case '%s', unexpected error: %s ", i, c.description, err)
			continue
		}
	}
}
