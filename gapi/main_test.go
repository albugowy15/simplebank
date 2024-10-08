package gapi

import (
	"testing"
	"time"

	db "github.com/albugowy15/simplebank/db/sqlc"
	"github.com/albugowy15/simplebank/utils"
	"github.com/albugowy15/simplebank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
