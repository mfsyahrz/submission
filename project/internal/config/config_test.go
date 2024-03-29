package config_test

import (
	"testing"

	. "github.com/mfsyahrz/submission/project/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	t.Run("failed init config - file not found", func(t *testing.T) {
		newConfig, err := New("./random_path")
		assert.NotNil(t, err)
		assert.Nil(t, newConfig)
	})

	t.Run("failed init config - invalid file", func(t *testing.T) {
		newConfig, err := New("../../test/config/env_invalid.test")
		assert.NotNil(t, err)
		assert.Nil(t, newConfig)
	})

	t.Run("success init config", func(t *testing.T) {
		wantConfig := &Config{
			Service: Service{
				Name: "sb_test",
				Port: Port{
					REST: "8080",
					GRPC: "8081",
				},
			},
			Postgres: Postgres{
				User:            "postgre_user",
				Password:        "postgre_pass",
				Name:            "postgre_name",
				Port:            "postgre_port",
				Host:            "postgre_host",
				MaxOpenConns:    50,
				MaxConnLifetime: 10,
				MaxIdleLifetime: 5,
			},
			OMDB: OMDB{
				Host:   "https://www.omdbapi.com",
				APIKey: "faf7e5bb",
			},
		}

		newConfig, err := New("../../test/config/env.test")
		assert.Nil(t, err)
		assert.Equal(t, wantConfig, newConfig)
	})
}

func TestRestandGRPCAddress(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		cfg, err := New("../../test/config/env.test")
		assert.NoError(t, err)

		cfg.Service.RestAddress()
		cfg.Service.GrpcAddress()
	})
}
