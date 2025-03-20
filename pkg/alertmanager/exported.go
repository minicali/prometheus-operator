package alertmanager

// ConfigBuilder is an exported type that wraps the internal alertmanagerConfig builder.
type ConfigBuilder struct {
	cfg       *alertmanagerConfig
	logger    *slog.Logger
	amVersion semver.Version
	store     *assets.StoreBuilder
	enforcer  enforcer
}

// NewConfigBuilder creates a new ConfigBuilder. 
// This function wraps the existing newConfigBuilder (which was unexported) and returns an exported type.
func NewConfigBuilder(logger *slog.Logger, amVersion semver.Version, store *assets.StoreBuilder, matcherStrategy monitoringv1.AlertmanagerConfigMatcherStrategy) *ConfigBuilder {
	// Call the internal function and cast the result to *ConfigBuilder.
	return (*ConfigBuilder)(newConfigBuilder(logger, amVersion, store, matcherStrategy))
}

// InitializeFromAlertmanagerConfig converts the given AlertmanagerConfig CR (and an optional global config)
// into the internal Alertmanager configuration. It wraps the internal initializeFromAlertmanagerConfig.
func (cb *ConfigBuilder) InitializeFromAlertmanagerConfig(ctx context.Context, globalConfig *monitoringv1.AlertmanagerGlobalConfig, amConfig *monitoringv1alpha1.AlertmanagerConfig) error {
	return cb.initializeFromAlertmanagerConfig(ctx, globalConfig, amConfig)
}

// Config returns the generated Alertmanager configuration.
func (cb *ConfigBuilder) Config() *alertmanagerConfig {
	return cb.cfg
}