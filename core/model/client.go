package model

import (
	"fmt"
	"time"

	"github.com/rollout/rox-go/core/context"
)

type BUID interface {
	fmt.Stringer
	GetValue() string
	GetQueryStringParts() map[string]string
}

type DeviceProperties interface {
	GetAllProperties() map[string]string
	RolloutEnvironment() string
	LibVersion() string
	DistinctID() string
	RolloutKey() string
}

type SelfManagedOptions interface {
	ServerURL() string
	AnalyticsURL() string
}

type RoxOptions interface {
	DevModeKey() string
	Version() string
	FetchInterval() time.Duration
	ImpressionHandler() ImpressionHandler
	ConfigurationFetchedHandler() ConfigurationFetchedHandler
	RoxyURL() string
	SelfManagedOptions() SelfManagedOptions
	DynamicPropertyRuleHandler() DynamicPropertyRuleHandler
}

type SdkSettings interface {
	APIKey() string
	DevModeSecret() string
}

type InternalFlags interface {
	IsEnabled(flagName string) bool
}

type DynamicAPI interface {
	IsEnabled(name string, defaultValue bool, ctx context.Context) bool
	StringValue(name string, defaultValue string, options []string, ctx context.Context) string
	IntValue(name string, defaultValue int, options []int, ctx context.Context) int
	DoubleValue(name string, defaultValue float64, options []float64, ctx context.Context) float64
}

type DynamicPropertyRuleHandler = func(args DynamicPropertyRuleHandlerArgs) interface{}

type DynamicPropertyRuleHandlerArgs struct {
	propName string
	context context.Context
}
