package app

import (
	"go-clean-template-v1/config"
	"go-clean-template-v1/pkg/postgres"

	"github.com/sulton0011/errs"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) (err error) {
	defer errs.WrapLog(&err, nil, "app - Run")
	// l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.Postgres.Url, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		return errs.Wrap(&err, "postgres.New")
	}
	defer pg.Close()

	// Use case
	// translationUseCase := usecase.New(
	// 	repo.New(pg),
	// 	webapi.New(),
	// )
	return nil
}
