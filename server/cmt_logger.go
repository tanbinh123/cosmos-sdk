package server

import (
	"cosmossdk.io/log"
	cmtlog "github.com/cometbft/cometbft/libs/log"
)

var _ cmtlog.Logger = (*CometZerologWrapper)(nil)

// CometZerologWrapper provides a wrapper around a zerolog.Logger instance.
// It implements CometBFT's Logger interface.
type CometZerologWrapper struct {
	log.ZeroLogWrapper
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (cmt CometZerologWrapper) With(keyVals ...interface{}) cmtlog.Logger {
	return cmt.With(keyVals...)
}
