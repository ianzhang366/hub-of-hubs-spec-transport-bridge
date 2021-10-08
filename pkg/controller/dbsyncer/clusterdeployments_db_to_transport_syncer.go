package dbsyncer

import (
	"fmt"
	"time"

	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/bundle"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/db"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/transport"
	cdv1 "github.com/openshift/hive/apis/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	clusterDeploymentTableName = "clusterdeployments"
	clusterDeploymentMsgKey    = "clusterdeployments"
)

// AddClusterDeploymentsDBToTransportSyncer adds clusterdeployments db to transport syncer to the manager.
func AddClusterDeploymentsDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName("clusterdeployments-db-to-transport-syncer"),
		db:                 db,
		dbTableName:        clusterDeploymentTableName,
		transport:          transport,
		transportBundleKey: clusterDeploymentMsgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &cdv1.ClusterDeployment{} },
		createBundleFunc:   bundle.NewBaseBundle,
	}); err != nil {
		return fmt.Errorf("failed to add db to transport syncer - %w", err)
	}

	return nil
}
