package featureflag

import (
	"github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/context"
)

type Option func(ctx *context.Context)

func WithUserID(userID string) Option {
	return func(ctx *context.Context) {
		ctx.UserId = userID
	}
}

func WithSessionID(sessionID string) Option {
	return func(ctx *context.Context) {
		ctx.SessionId = sessionID
	}
}

type FeatureFlag struct {
	client *unleash.Client
}

func (f FeatureFlag) IsEnabled(featureName string, options ...Option) bool {
	ctx := context.Context{}

	for _, option := range options {
		option(&ctx)
	}

	return f.client.IsEnabled(featureName, unleash.WithContext(ctx))
}

func (f *FeatureFlag) IsReady() {
	// todo
}

func (f *FeatureFlag) Close() error {
	return f.client.Close()
}

func NewFeatureFlag(
	environment string, // The name of the environment the application runs in (not the name of the application itself).
	apiURL string, // URL where the client (application) connects to get a list of feature flags.
	instanceID string,
) (*FeatureFlag, error) {
	client, err := unleash.NewClient(
		unleash.WithAppName(environment),
		unleash.WithUrl(apiURL),
		unleash.WithInstanceId(instanceID),
		unleash.WithListener(&unleash.DebugListener{}),
		// This client uses go routines to report several events and doesn't drain the channel by default.
		// So you need to either register a listener using WithListener or drain the channel "manually"
	)

	return &FeatureFlag{
		client: client,
	}, err
}

//The Go client comes with implementations for the built-in activation strategies provided by unleash.
//
//	DefaultStrategy
//UserIdStrategy
//FlexibleRolloutStrategy
//GradualRolloutUserIdStrategy
//GradualRolloutSessionIdStrategy
//GradualRolloutRandomStrategy
//RemoteAddressStrategy
//ApplicationHostnameStrategy
