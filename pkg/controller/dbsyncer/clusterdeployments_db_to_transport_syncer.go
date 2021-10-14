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
	clusterDeploymentMsgKey    = "ClusterDeployments"
)

// AddClusterDeploymentsDBToTransportSyncer adds clusterdeployments db to transport syncer to the manager.
func AddClusterDeploymentsDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	componentName := "clusterdeployments-db-to-transport-syncer"
	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName(componentName),
		db:                 db,
		dbTableName:        clusterDeploymentTableName,
		transport:          transport,
		transportBundleKey: clusterDeploymentMsgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &cdv1.ClusterDeployment{} },
		createBundleFunc:   bundle.NewClusterLifecycleBundle,
	}); err != nil {
		return fmt.Errorf("failed to add %s db to transport syncer - %w", componentName, err)
	}

	return nil
}
