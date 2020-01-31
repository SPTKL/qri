package remote

import (
	"context"

	"github.com/qri-io/qri/dsref"
	"github.com/qri-io/qri/logbook/oplog"
	reporef "github.com/qri-io/qri/repo/ref"
)

// Client connects to remotes to perform synchronization tasks
type Client interface {
	ListDatasets(ctx context.Context, ds *reporef.DatasetRef, term string, offset, limit int) ([]reporef.DatasetRef, error)

	PushDataset(ctx context.Context, ref reporef.DatasetRef, remoteAddr string) error
	PullDataset(ctx context.Context, ref *reporef.DatasetRef, remoteAddr string) error
	RemoveDataset(ctx context.Context, ref reporef.DatasetRef, remoteAddr string) error
	AddDataset(ctx context.Context, ref *reporef.DatasetRef, remoteAddr string) error

	PushLogs(ctx context.Context, ref dsref.Ref, remoteAddr string) error
	FetchLogs(ctx context.Context, ref dsref.Ref, remoteAddr string) (*oplog.Log, error)
	CloneLogs(ctx context.Context, ref dsref.Ref, remoteAddr string) error
	RemoveLogs(ctx context.Context, ref dsref.Ref, remoteAddr string) error

	ResolveHeadRef(ctx context.Context, ref *reporef.DatasetRef, remoteAddr string) error
}
