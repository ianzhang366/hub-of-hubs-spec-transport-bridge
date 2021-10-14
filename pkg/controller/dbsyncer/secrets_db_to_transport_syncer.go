package dbsyncer

import (
	"fmt"
	"time"

	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/bundle"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/db"
	"github.com/open-cluster-management/hub-of-hubs-spec-transport-bridge/pkg/transport"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	secretsSpecTableName = "secrets"
	SecretMsgKey         = "Secrets"
)

// AddSecretsDBToTransportSyncer adds secrets db to transport syncer to the manager.
func AddSecretsDBToTransportSyncer(mgr ctrl.Manager, db db.HubOfHubsSpecDB, transport transport.Transport,
	syncInterval time.Duration) error {
	componentName := "secrets-db-transport-syncer"
	if err := mgr.Add(&genericDBToTransportSyncer{
		log:                ctrl.Log.WithName(componentName),
		db:                 db,
		dbTableName:        secretsSpecTableName,
		transport:          transport,
		transportBundleKey: SecretMsgKey,
		syncInterval:       syncInterval,
		createObjFunc:      func() metav1.Object { return &corev1.Secret{} },
		createBundleFunc:   bundle.NewClusterLifecycleBundle,
	}); err != nil {
		return fmt.Errorf("failed to add %s db to transport syncer - %w", componentName, err)
	}

	return nil
}
