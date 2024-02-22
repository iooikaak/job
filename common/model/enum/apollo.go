package enum

type Apollo string

const (
	ApolloNamespaces    Apollo = "APOLLO_NAMESPACES"
	ApolloAppID         Apollo = "APOLLO_APP_ID"
	ApolloCluster       Apollo = "APOLLO_CLUSTER"
	ApolloClusterValue  Apollo = "Dev"
	ApolloMetaAddr      Apollo = "APOLLO_META_ADDR"
	ApolloMetaAddrValue Apollo = "http://cyril.develop:8080"
	ApolloCacheDir      Apollo = "APOLLO_CACHE_DIR"
	ApolloCacheDirValue Apollo = "/tmp"
	ApolloAppName       Apollo = "TEST1.job.json"
)

func (a Apollo) String() string {
	return string(a)
}
