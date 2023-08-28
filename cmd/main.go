package main

import (
	"go-clean-template-v1/config"

	"github.com/sulton0011/errs"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		errs.WrapLog(&err, nil, "main")
		return
	}

	println(cfg.Postgres.Url)

}
