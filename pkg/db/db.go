package db

import (
	"context"
	"time"

	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/bundle"
)

// HubOfHubsSpecDB is the needed interface for the db transport bridge.
type HubOfHubsSpecDB interface {
	GetBundle(ctx context.Context, tableName string, createObjFunc bundle.CreateObjectFunction,
		intoBundle bundle.Bundle) (*time.Time, error)
	GetLastUpdateTimestamp(ctx context.Context, tableName string) (*time.Time, error)
	Stop()
}
