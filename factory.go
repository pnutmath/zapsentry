package zapsentry

import (
	"github.com/getsentry/sentry-go"
)

// NewSentryClientFromDSN creates new sentry client from DSN value
func NewSentryClientFromDSN(DSN string) SentryClientFactory {
	return func() (*sentry.Client, error) {
		return sentry.NewClient(sentry.ClientOptions{
			Dsn: DSN,
		})
	}
}

// NewSentryClientFromClient creates new sentry client from sentry client reference
func NewSentryClientFromClient(client *sentry.Client) SentryClientFactory {
	return func() (*sentry.Client, error) {
		return client, nil
	}
}

// SentryClientFactory function that return sentry client
type SentryClientFactory func() (*sentry.Client, error)
