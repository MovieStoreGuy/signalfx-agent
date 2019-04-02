package sql

import (
	"database/sql"

	"github.com/apex/log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/signalfx/signalfx-agent/internal/core/config"
	"github.com/signalfx/signalfx-agent/internal/monitors"
)

const monitorType = "sql"

var logger = log.WithFields(log.Fields{"monitorType": monitorType})

func init() {
	monitors.Register(monitorType, func() interface{} { return &Monitor{} }, &Config{})
}

type sqlQuery struct {
}

// Config for this monitor
type Config struct {
	config.MonitorConfig `yaml:",inline" acceptsEndpoints:"true"`

	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`

	DBDriver string `yaml:"dbDriver"`

	// A list of queries to make against the database that are used to generate
	// datapoints.
	Queries []sqlQuery `yaml:"queries"`
}

type Monitor struct {
	database *sql.DB
}

func (m *Monitor) Shutdown() {
	if m.database != nil {
		m.database.Close()
	}
}
