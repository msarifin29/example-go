package test

import (
	"gorm-example/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViperConfig(t *testing.T) {
	config := config.NewViper()
	err := config.ReadInConfig()
	assert.Nil(t, err)
}
