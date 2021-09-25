package datasource

import (
	"errors"
	"fmt"

	"github.com/MonsterYNH/nava/datasource/prometheus"
	"github.com/MonsterYNH/nava/setting"

	"github.com/go-pg/pg/v10"
)

type DataSource interface {
	PostgresDB() (*pg.DB, error)
	PrometheusClient() (*prometheus.Client, error)
}

func InitDataSource(config *setting.Config) (DataSource, error) {
	datasource := &DataSourceEntry{config: *config}
	if config.DataSource.Postgres.Enable {
		if err := datasource.initPostgresDB(config.DataSource.Postgres); err != nil {
			return nil, err
		}
	}

	if config.DataSource.Promethes.Enable {
		if err := datasource.initPrometheusClient(config.DataSource.Promethes); err != nil {
			return nil, err
		}
	}

	return datasource, nil
}

type DataSourceEntry struct {
	config           setting.Config
	postgresDB       *pg.DB
	prometheusClient *prometheus.Client
}

func (datasource *DataSourceEntry) PostgresDB() (*pg.DB, error) {
	if !datasource.config.DataSource.Postgres.Enable {
		return nil, errors.New("datasource postgres is not enable")
	}

	if datasource.postgresDB == nil {
		return nil, errors.New("datasource postgres is not init")
	}

	return datasource.postgresDB, nil
}

func (datasource *DataSourceEntry) PrometheusClient() (*prometheus.Client, error) {
	if !datasource.config.DataSource.Promethes.Enable {
		return nil, errors.New("datasource prometheus api is not enable")
	}

	if datasource.prometheusClient == nil {
		return nil, errors.New("datasource prometheus is not init")
	}

	return datasource.prometheusClient, nil
}

func (dataSource *DataSourceEntry) initPostgresDB(config setting.PostgresConfig) error {
	if !config.Enable {
		return fmt.Errorf("datasource postgres is not enable")
	}

	opt, err := pg.ParseURL(config.URI)
	if err != nil {
		return err
	}

	dataSource.postgresDB = pg.Connect(opt)
	return nil
}

func (dataSource *DataSourceEntry) initPrometheusClient(config setting.PrometheusConfig) error {
	if !config.Enable {
		return fmt.Errorf("datasource prometheus api is not enable")
	}

	client, err := prometheus.NewClient(config.URL)
	dataSource.prometheusClient = client
	return err
}
