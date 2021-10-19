package dbsyncer

import (
	"fmt"
	"time"

	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/bundle"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/db"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/transport"
	agentv1 "github.com/open-cluster-management/klusterlet-addon-controller/pkg/apis/agent/v1"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// AddClusterDeploymentsDBToTransportSyncer adds clusterdeployments db to transport syncer to the manager.
func AddClusterDeploymentsDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	component := "clusterdeployments"
	msgKey := "ClusterDeployments"

	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName(fmt.Sprintf("%s-db-to-transport-syncer", component)),
		db:                 db,
		dbTableName:        component,
		transport:          transport,
		transportBundleKey: msgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &hivev1.ClusterDeployment{} },
		createBundleFunc:   bundle.NewClusterLifecycleBundle,
	}); err != nil {
		return fmt.Errorf("failed to add %s db to transport syncer - %w", component, err)
	}

	return nil
}

func AddMachinepoolDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	component := "machinepools"
	msgKey := "Machinepools"

	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName(fmt.Sprintf("%s-db-to-transport-syncer", component)),
		db:                 db,
		dbTableName:        component,
		transport:          transport,
		transportBundleKey: msgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &hivev1.ClusterDeployment{} },
		createBundleFunc:   bundle.NewClusterLifecycleBundle,
	}); err != nil {
		return fmt.Errorf("failed to add %s db to transport syncer - %w", component, err)
	}

	return nil
}

func AddKlusterletaddonconfigDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	component := "klusterletaddonconfigs"
	msgKey := "Klusterletaddonconfigs"

	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName(fmt.Sprintf("%s-db-to-transport-syncer", component)),
		db:                 db,
		dbTableName:        component,
		transport:          transport,
		transportBundleKey: msgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &agentv1.KlusterletAddonConfig{} },
		createBundleFunc:   bundle.NewClusterLifecycleBundle,
	}); err != nil {
		return fmt.Errorf("failed to add %s db to transport syncer - %w", component, err)
	}

	return nil
}
