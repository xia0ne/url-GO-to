package model

import (
	"context"
)

var (
	MyConfigs = &Config{}
	MyRedis   = &StorageService{}
	Ctx       = context.Background()
	MyStore   = &URLStore{}
)
