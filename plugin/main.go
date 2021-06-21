package main

import (
	"context"

	config2 "github.com/nori-plugins/profile/internal/config"

	"github.com/nori-plugins/profile/internal/domain/service"
	"gorm.io/gorm"

	"github.com/nori-io/common/v5/pkg/domain/config"
	em "github.com/nori-io/common/v5/pkg/domain/enum/meta"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-io/common/v5/pkg/domain/meta"
	p "github.com/nori-io/common/v5/pkg/domain/plugin"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	m "github.com/nori-io/common/v5/pkg/meta"
	noriGorm "github.com/nori-io/interfaces/database/gorm"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	instance service.ProfileService
	config   config2.Config
	logger   logger.FieldLogger
}

func (p *plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = config2.Config{}
	p.logger = log

	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p *plugin) Meta() meta.Meta {
	return m.Meta{
		ID: m.ID{
			ID:      "nori/profile.go/Profile",
			Version: "0.1.0",
		},
		Author: m.Author{
			Name: "Nori Authors",
			URL:  "https://nori.io",
		},
		Dependencies: nil,
		Description: m.Description{
			Title:       "Nori HTTP Interface",
			Description: "Official implementation of nori/http/HTTP interface",
		},
		Interface: meta.NewInterface("", ""),
		License: []meta.License{m.License{
			Title: "",
			Type:  em.Apache2_0,
			URL:   "http://www.apache.org/licenses/",
		}},
		Links: nil,
		Repository: m.Repository{
			Type: em.Git,
			URL:  "github.com/nori-plugins/http",
		},
		Tags: []string{"nori", "profile.go"},
	}
}

func (p *plugin) Start(ctx context.Context, registry registry.Registry) error {
	config := p.config

	_, err := Initialize(registry, config, p.logger)
	return err
}

func (p *plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Install(_ context.Context, registry registry.Registry) error {
	db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}
	//@todo actual sql code
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(`CREATE TABLE profiles(
	  	id bigserial PRIMARY KEY,
		user_id bigint NOT NULL,
		first_name  VARCHAR (25) NOT NULL,
		last_name VARCHAR (25) NOT NULL,
		nickname   VARCHAR (25) NOT NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
			);
		`).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (p plugin) UnInstall(_ context.Context, registry registry.Registry) error {
	db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(`drop table profiles;
		`).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
