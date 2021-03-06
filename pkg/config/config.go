package config

import (
	"time"

	"github.com/appscode/go/log/golog"
)

var (
	AnalyticsClientID string
	EnableAnalytics   = true
	LoggerOptions     golog.Options
	BuiltinTemplates  = "/srv/voyager/templates/*.cfg"
)

type Config struct {
	Burst                       int
	CloudConfigFile             string
	CloudProvider               string
	EnableRBAC                  bool
	HAProxyImage                string
	ExporterImage               string
	IngressClass                string
	MaxNumRequeues              int
	NumThreads                  int
	OperatorNamespace           string
	OperatorService             string
	OpsAddress                  string
	QPS                         float32
	RestrictToOperatorNamespace bool
	ResyncPeriod                time.Duration
	WatchNamespace              string
}
