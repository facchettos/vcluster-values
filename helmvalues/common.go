package helmvalues

type MonitoringValues struct {
	ServiceMonitor ServiceMonitor `json:"serviceMonitor,omitempty"`
}

type ServiceMonitor struct {
	Enabled bool `json:"enabled,omitempty"`
}

type EmbeddedEtcdValues struct {
	Enabled         bool `json:"enabled,omitempty"`
	MigrateFromEtcd bool `json:"migrateFromEtcd,omitempty"`
}

type Storage struct {
	Persistence bool   `json:"persistence,omitempty"`
	Size        string `json:"size,omitempty"`
}

type BaseHelm struct {
	GlobalAnnotations    map[string]string        `json:"globalAnnotations,omitempty"`
	Pro                  bool                     `json:"pro,omitempty"`
	Headless             bool                     `json:"headless,omitempty"`
	DefaultImageRegistry string                   `json:"defaultImageRegistry,omitempty"`
	Plugin               map[string]interface{}   `json:"plugin,omitempty"`
	Sync                 SyncValues               `json:"sync,omitempty"`
	FallbackHostDNS      bool                     `json:"fallbackHostDns,omitempty"`
	MapServices          MapServices              `json:"mapServices,omitempty"`
	Proxy                ProxyValues              `json:"proxy,omitempty"`
	Volumes              []map[string]interface{} `json:"volumes,omitempty"`
	ServiceAccount       struct {
		Create bool `json:"create,omitempty"`
	} `json:"serviceAccount,omitempty"`
	WorkloadServiceAccount struct {
		Annotations map[string]string `json:"annotations,omitempty"`
	} `json:"workloadServiceAccount,omitempty"`
	Rbac                RBACValues               `json:"rbac,omitempty"`
	NodeSelector        map[string]interface{}   `json:"nodeSelector,omitempty"`
	Affinity            map[string]interface{}   `json:"affinity,omitempty"`
	PriorityClassName   string                   `json:"priorityClassName,omitempty"`
	Tolerations         []map[string]interface{} `json:"tolerations,omitempty"`
	Labels              map[string]string        `json:"labels,omitempty"`
	PodLabels           map[string]string        `json:"podLabels,omitempty"`
	Annotations         map[string]string        `json:"annotations,omitempty"`
	PodAnnotations      map[string]string        `json:"podAnnotations,omitempty"`
	PodDisruptionBudget PDBValues                `json:"podDisruptionBudget,omitempty"`
	Service             ServiceValues            `json:"service,omitempty"`
	Ingress             IngressValues            `json:"ingress,omitempty"`

	SecurityContext    map[string]interface{} `json:"securityContext,omitempty"`
	PodSecurityContext map[string]interface{} `json:"podSecurityContext,omitempty"`
	Openshift          struct {
		Enable bool `json:"enable,omitempty"`
	} `json:"openshift,omitempty"`
	Coredns            CoreDNSValues    `json:"coredns,omitempty"`
	Isolation          IsolationValues  `json:"isolation,omitempty"`
	Init               InitValues       `json:"init,omitempty"`
	MultiNamespaceMode EnabledSwitch    `json:"multiNamespaceMode,omitempty"`
	Telemetry          TelemetryValues  `json:"telemetry,omitempty"`
	NoopSyncer         NoopSyncerValues `json:"noopSyncer,omitempty"`
	Monitoring         MonitoringValues `json:"monitoring,omitempty"`
	Admission          AdmissionValues  `json:"admission,omitempty"`
}

type SyncerValues struct {
	ControlPlaneCommonValues
	ExtraArgs             []string                 `json:"extraArgs,omitempty"`
	Env                   []map[string]interface{} `json:"env,omitempty"`
	LivenessProbe         EnabledSwitch            `json:"livenessProbe,omitempty"`
	ReadinessProbe        EnabledSwitch            `json:"readinessProbe,omitempty"`
	VolumeMounts          []map[string]interface{} `json:"volumeMounts,omitempty"`
	ExtraVolumeMounts     []map[string]interface{} `json:"extraVolumeMounts,omitempty"`
	Resources             map[string]interface{}   `json:"resources,omitempty"`
	KubeConfigContextName string                   `json:"kubeConfigContextName,omitempty"`
	ServiceAnnotations    map[string]string        `json:"serviceAnnotations,omitempty"`
	Replicas              uint32                   `json:"replicas,omitempty"`
	Storage               Storage                  `json:"storage,omitempty"`
	Labels                map[string]string        `json:"labels,omitempty"`
	Annotations           map[string]string        `json:"annotations,omitempty"`
}

type SyncValues struct {
	Services               EnabledSwitch  `json:"services,omitempty"`
	Configmaps             SyncConfigMaps `json:"configmaps,omitempty"`
	Secrets                SyncSecrets    `json:"secrets,omitempty"`
	Endpoints              EnabledSwitch  `json:"endpoints,omitempty"`
	Pods                   SyncPods       `json:"pods,omitempty"`
	Events                 EnabledSwitch  `json:"events,omitempty"`
	PersistentVolumeClaims EnabledSwitch  `json:"persistentVolumeClaims,omitempty"`
	Ingresses              EnabledSwitch  `json:"ingresses,omitempty"`
	Ingressclasses         EnabledSwitch  `json:"ingressclasses,omitempty"`
	FakeNodes              EnabledSwitch  `json:"fake-nodes,omitempty"`
	FakePersistentvolumes  EnabledSwitch  `json:"fake-persistentvolumes,omitempty"`
	Nodes                  SyncNodes      `json:"nodes,omitempty"`
	PersistentVolumes      EnabledSwitch  `json:"persistentVolumes,omitempty"`
	StorageClasses         EnabledSwitch  `json:"storageClasses,omitempty"`
	Hoststorageclasses     EnabledSwitch  `json:"hoststorageclasses,omitempty"`
	Priorityclasses        EnabledSwitch  `json:"priorityclasses,omitempty"`
	Networkpolicies        EnabledSwitch  `json:"networkpolicies,omitempty"`
	Volumesnapshots        EnabledSwitch  `json:"volumesnapshots,omitempty"`
	Poddisruptionbudgets   EnabledSwitch  `json:"poddisruptionbudgets,omitempty"`
	Serviceaccounts        EnabledSwitch  `json:"serviceaccounts,omitempty"`
	Generic                SyncGeneric    `json:"generic,omitempty"`
}

type SyncConfigMaps struct {
	Enabled bool `json:"enabled,omitempty"`
	All     bool `json:"all,omitempty"`
}

type SyncSecrets struct {
	Enabled bool `json:"enabled,omitempty"`
	All     bool `json:"all,omitempty"`
}

type SyncPods struct {
	Enabled             bool `json:"enabled,omitempty"`
	EphemeralContainers bool `json:"ephemeralContainers,omitempty"`
	Status              bool `json:"status,omitempty"`
}

type SyncNodes struct {
	FakeKubeletIPs  bool   `json:"fakeKubeletIPs,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	SyncAllNodes    bool   `json:"syncAllNodes,omitempty"`
	NodeSelector    string `json:"nodeSelector,omitempty"`
	EnableScheduler bool   `json:"enableScheduler,omitempty"`

	// Deprecated: should be removed from the chart first
	SyncNodeChanges bool `json:"syncNodeChanges,omitempty"`
}

type SyncGeneric struct {
	Config string `json:"config,omitempty"`
}

type EnabledSwitch struct {
	Enabled bool `json:"enabled,omitempty"`
}

type MapServices struct {
	FromVirtual []ServiceMapping `json:"fromVirtual,omitempty"`
	FromHost    []ServiceMapping `json:"fromHost,omitempty"`
}

type ServiceMapping struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ProxyValues struct {
	MetricsServer MetricsProxyServerConfig `json:"metricsServer,omitempty"`
}

type MetricsProxyServerConfig struct {
	Nodes EnabledSwitch `json:"nodes,omitempty"`
	Pods  EnabledSwitch `json:"pods,omitempty"`
}

type VClusterValues struct {
	Image             string                   `json:"image,omitempty"`
	ImagePullPolicy   string                   `json:"imagePullPolicy,omitempty"`
	Command           []string                 `json:"command,omitempty"`
	BaseArgs          []string                 `json:"baseArgs,omitempty"`
	ExtraArgs         []string                 `json:"extraArgs,omitempty"`
	ExtraVolumeMounts []map[string]interface{} `json:"extraVolumeMounts,omitempty"`
	VolumeMounts      []map[string]interface{} `json:"volumeMounts,omitempty"`
	Env               []map[string]interface{} `json:"env,omitempty"`
	Resources         map[string]interface{}   `json:"resources,omitempty"`

	// this is only provided in context of k0s right now
	PriorityClassName string `json:"priorityClassName,omitempty"`
}

// These should be remove from the chart first as they are deprecated there
type RBACValues struct {
	ClusterRole struct {
		Create bool `json:"create,omitempty"`
	} `json:"clusterRole,omitempty"`
	Role struct {
		Create               bool     `json:"create,omitempty"`
		Extended             bool     `json:"extended,omitempty"`
		ExcludedAPIResources []string `json:"excludedAPIResources,omitempty"`
	} `json:"role,omitempty"`
}

type PDBValues struct {
	Enabled        bool        `json:"enabled,omitempty"`
	MinAvailable   interface{} `json:"minAvailable,omitempty"`
	MaxUnavailable interface{} `json:"maxUnavailable,omitempty"`
}

type ServiceValues struct {
	Type                     string            `json:"type,omitempty"`
	ExternalIPs              []string          `json:"externalIPs,omitempty"`
	ExternalTrafficPolicy    string            `json:"externalTrafficPolicy,omitempty"`
	LoadBalancerIP           string            `json:"loadBalancerIP,omitempty"`
	LoadBalancerSourceRanges []string          `json:"loadBalancerSourceRanges,omitempty"`
	LoadBalancerClass        string            `json:"loadBalancerClass,omitempty"`
	LoadBalancerAnnotation   map[string]string `json:"loadBalancerAnnotations,omitempty"`
}

type IngressValues struct {
	Enabled          bool                     `json:"enabled,omitempty"`
	PathType         string                   `json:"pathType,omitempty"`
	IngressClassName string                   `json:"ingressClassName,omitempty"`
	Host             string                   `json:"host,omitempty"`
	Annotations      map[string]string        `json:"annotations,omitempty"`
	TLS              []map[string]interface{} `json:"tls,omitempty"`
}

type CoreDNSValues struct {
	Integrated     bool                   `json:"integrated,omitempty"`
	Plugin         CoreDNSPluginValues    `json:"plugin,omitempty"`
	Enabled        bool                   `json:"enabled,omitempty"`
	Replicas       uint32                 `json:"replicas,omitempty"`
	NodeSelector   map[string]interface{} `json:"nodeSelector,omitempty"`
	Image          string                 `json:"image,omitempty"`
	Config         string                 `json:"config,omitempty"`
	Service        CoreDNSServiceValues   `json:"service,omitempty"`
	Resources      map[string]interface{} `json:"resources,omitempty"`
	Manifests      string                 `json:"manifests,omitempty"`
	PodAnnotations map[string]string      `json:"podAnnotations,omitempty"`
	PodLabels      map[string]string      `json:"podLabels,omitempty"`
}

type CoreDNSPluginValues struct {
	Enabled bool          `json:"enabled,omitempty"`
	Config  []DNSMappings `json:"config,omitempty"`
}

type DNSMappings struct {
	Record    Record       `json:"record,omitempty"`
	Target    Target       `json:"target,omitempty"`
	AllowedOn []FilterSpec `json:"allowedOn,omitempty"`
	ExceptOn  []FilterSpec `json:"exceptOn,omitempty"`
}

type Record struct {
	RecordType RecordType `json:"recordType,omitempty"`
	FQDN       *string    `json:"fqdn,omitempty"`
	Service    *string    `json:"service,omitempty"`
	Namespace  *string    `json:"namespace,omitempty"`
}

type RecordType string
type TargetMode string

type Target struct {
	Mode      TargetMode `json:"mode,omitempty"`
	VCluster  *string    `json:"vcluster,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Service   *string    `json:"service,omitempty"`
	Namespace *string    `json:"namespace,omitempty"`
}

type FilterSpec struct {
	Name      string   `json:"name,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type CoreDNSServiceValues struct {
	Type                  string            `json:"type,omitempty"`
	ExternalIPs           []string          `json:"externalIPs,omitempty"`
	ExternalTrafficPolicy string            `json:"externalTrafficPolicy,omitempty"`
	Annotations           map[string]string `json:"annotations,omitempty"`
}

type IsolationValues struct {
	Enabled             bool          `json:"enabled,omitempty"`
	Namespace           *string       `json:"namespace,omitempty"`
	PodSecurityStandard string        `json:"podSecurityStandard,omitempty"`
	NodeProxyPermission EnabledSwitch `json:"nodeProxyPermission,omitempty"`

	ResourceQuota struct {
		Enabled       bool                     `json:"enabled,omitempty"`
		Quota         map[string]interface{}   `json:"quota,omitempty"`
		ScopeSelector map[string]interface{}   `json:"scopeSelector,omitempty"`
		Scopes        []map[string]interface{} `json:"scopes,omitempty"`
	} `json:"resourceQuota,omitempty"`

	LimitRange    IsolationLimitRangeValues `json:"limitRange,omitempty"`
	NetworkPolicy NetworkPolicyValues       `json:"networkPolicy,omitempty"`
}

type IsolationLimitRangeValues struct {
	Enabled        bool                             `json:"enabled,omitempty"`
	Default        IsolationLimitRangeDefaultValues `json:"default,omitempty"`
	DefaultRequest IsolationLimitRangeDefaultValues `json:"defaultRequest,omitempty"`
}

type IsolationLimitRangeDefaultValues struct {
	EphemeralStorage string `json:"ephemeral-storage,omitempty"`
	Memory           string `json:"memory,omitempty"`
	CPU              string `json:"cpu,omitempty"`
}

type NetworkPolicyValues struct {
	Enabled             bool `json:"enabled,omitempty"`
	OutgoingConnections struct {
		IPBlock struct {
			CIDR   string   `json:"cidr,omitempty"`
			Except []string `json:"except,omitempty"`
		} `json:"ipBlock,omitempty"`
	} `json:"outgoingConnections,omitempty"`
}

type InitValues struct {
	Manifests         string           `json:"manifests,omitempty"`
	ManifestsTemplate string           `json:"manifestsTemplate,omitempty"`
	Helm              []InitHelmCharts `json:"helm,omitempty"`
}

type InitHelmCharts struct {
	Bundle string `json:"bundle,omitempty"`
	Chart  struct {
		Name     string `json:"name,omitempty"`
		Version  string `json:"version,omitempty"`
		Repo     string `json:"repo,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Insecure bool   `json:"insecure,omitempty"`
	} `json:"chart,omitempty"`
	Release struct {
		ReleaseName      string `json:"releaseName,omitempty"`
		ReleaseNamespace string `json:"releaseNamespace,omitempty"`
		Timeout          uint32 `json:"timeout,omitempty"`
	} `json:"release,omitempty"`
	Values         string `json:"values,omitempty"`
	ValuesTemplate string `json:"valuesTemplate,omitempty"`
}

type TelemetryValues struct {
	Disabled           bool   `json:"disabled,omitempty"`
	InstanceCreator    string `json:"instanceCreator,omitempty"`
	PlatformUserID     string `json:"platformUserID,omitempty"`
	PlatformInstanceID string `json:"platformInstanceID,omitempty"`
	MachineID          string `json:"machineID,omitempty"`
}

type NoopSyncerValues struct {
	Enabled        bool `json:"enabled,omitempty"`
	Synck8sService bool `json:"synck8SService,omitempty"`
	Secret         struct {
		ServerCaCert        string `json:"serverCaCert,omitempty"`
		ServerCaKey         string `json:"serverCaKey,omitempty"`
		ClientCaCert        string `json:"clientCaCert,omitempty"`
		RequestHeaderCaCert string `json:"requestHeaderCaCert,omitempty"`
		KubeConfig          string `json:"kubeConfig,omitempty"`
	}
}

type AdmissionValues struct {
	ValidatingWebhooks []string `json:"validatingWebhooks,omitempty"`
	MutatingWebhooks   []string `json:"mutatingWebhooks,omitempty"`
}

type ControlPlaneCommonValues struct {
	Image           string `json:"image,omitempty"`
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
}

type SyncerExORCommonValues struct {
	ExtraArgs []string               `json:"extraArgs,omitempty"`
	Resources map[string]interface{} `json:"resources,omitempty"`
}

type CommonValues struct {
	Volumes           []map[string]interface{} `json:"volumes,omitempty"`
	PriorityClassName string                   `json:"priorityClassName,omitempty"`
	NodeSelector      map[string]interface{}   `json:"nodeSelector,omitempty"`
	Affinity          map[string]interface{}   `json:"affinity,omitempty"`
	Tolerations       []map[string]interface{} `json:"tolerations,omitempty"`
	PodAnnotations    map[string]string        `json:"podAnnotations,omitempty"`
	PodLabels         map[string]string        `json:"podLabels,omitempty"`
}
