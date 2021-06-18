package main

import (
	"context"

	"github.com/nori-plugins/profile/internal/domain/service"

	"github.com/nori-io/common/v5/pkg/domain/config"
	em "github.com/nori-io/common/v5/pkg/domain/enum/meta"
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-io/common/v5/pkg/domain/meta"
	p "github.com/nori-io/common/v5/pkg/domain/plugin"
	"github.com/nori-io/common/v5/pkg/domain/registry"
	m "github.com/nori-io/common/v5/pkg/meta"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	instance service.ProfileService
	/*	config conf
		log    logger.FieldLogger*/
}

type conf struct {
}

func (p *plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	/*p.config = conf{
		host: config.String("host", "http server host"),
		port: config.String("port", "http server port"),
	}
	p.log = log*/

	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p *plugin) Meta() meta.Meta {
	return m.Meta{
		ID: m.ID{
			ID:      "nori/profile/Profile",
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
		Tags: []string{"nori", "profile"},
	}
}

func (p *plugin) Start(ctx context.Context, registry registry.Registry) error {
	/*addr := fmt.Sprintf("%s:%s", p.config.host(), p.config.port())
	p.log.Info(fmt.Sprintf("http addr %s", addr))
	p.server = server.New(addr)*/
	return nil
}

func (p *plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Install(_ context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) UnInstall(_ context.Context, registry registry.Registry) error {
	return nil
}
