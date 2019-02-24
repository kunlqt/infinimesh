package node

import (
	"context"

	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

const (
	KindAsset  = "asset"
	KindDevice = "device"
)

type Repo interface {
	CreateUserAccount(ctx context.Context, username, password string, isRoot bool) (uid string, err error)
	ListAccounts(ctx context.Context) (accounts []*nodepb.Account, err error)
	GetAccount(ctx context.Context, accountID string) (account *nodepb.Account, err error)

	IsAuthorized(ctx context.Context, target, who, action string) (decision bool, err error)
	IsAuthorizedNamespace(ctx context.Context, namespace, account string, action nodepb.Action) (decision bool, err error)
	Authorize(ctx context.Context, account, node, action string, inherit bool) (err error)
	AuthorizeNamespace(ctx context.Context, account, namespace string, action nodepb.Action) (err error)
	Authenticate(ctx context.Context, username, password string) (success bool, uid string, err error)

	CreateObject(ctx context.Context, name, parentID, kind, namespaceID string) (id string, err error)
	DeleteObject(ctx context.Context, uid string) (err error)
	ListForAccount(ctx context.Context, account string) (inheritedObjects []*nodepb.Object, err error)

	CreateNamespace(ctx context.Context, name string) (id string, err error)
	GetNamespace(ctx context.Context, uid string) (namespace *nodepb.Namespace, err error)
	ListNamespaces(ctx context.Context) (namespaces []*nodepb.Namespace, err error)
	ListNamespacesForAccount(ctx context.Context, accountID string) (namespaces []*nodepb.Namespace, err error)

	ListInNamespaceForAccount(ctx context.Context, accountID, namespaceName string, recurse bool) (objects []*nodepb.Object, err error)
}
