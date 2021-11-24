package config

import (
	"context"
	"path"

	"github.com/owncloud/ocis/ocis-pkg/shared"

	"github.com/owncloud/ocis/ocis-pkg/config/defaults"
)

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string `ocisConfig:"addr"`
	Token  string `ocisConfig:"token"`
	Pprof  bool   `ocisConfig:"pprof"`
	Zpages bool   `ocisConfig:"zpages"`
}

// Service defines the available service configuration.
type Service struct {
	Name    string `ocisConfig:"name"`
	Version string `ocisConfig:"version"`
}

// Tracing defines the available tracing configuration.
type Tracing struct {
	Enabled   bool   `ocisConfig:"enabled"`
	Type      string `ocisConfig:"type"`
	Endpoint  string `ocisConfig:"endpoint"`
	Collector string `ocisConfig:"collector"`
	Service   string `ocisConfig:"service"`
}

// Ldap defined the available LDAP configuration.
type Ldap struct {
	Enabled   bool   `ocisConfig:"enabled"`
	Addr      string `ocisConfig:"addr"`
	Namespace string `ocisConfig:"namespace"`
}

// Ldaps defined the available LDAPS configuration.
type Ldaps struct {
	Enabled   bool   `ocisConfig:"enabled"`
	Addr      string `ocisConfig:"addr"`
	Namespace string `ocisConfig:"namespace"`
	Cert      string `ocisConfig:"cert"`
	Key       string `ocisConfig:"key"`
}

// Backend defined the available backend configuration.
type Backend struct {
	Datastore   string   `ocisConfig:"datastore"`
	BaseDN      string   `ocisConfig:"base_dn"`
	Insecure    bool     `ocisConfig:"insecure"`
	NameFormat  string   `ocisConfig:"name_format"`
	GroupFormat string   `ocisConfig:"group_format"`
	Servers     []string `ocisConfig:"servers"`
	SSHKeyAttr  string   `ocisConfig:"ssh_key_attr"`
	UseGraphAPI bool     `ocisConfig:"use_graph_api"`
}

// Config combines all available configuration parts.
type Config struct {
	*shared.Commons

	File           string      `ocisConfig:"file"`
	Log            *shared.Log `ocisConfig:"log"`
	Debug          Debug       `ocisConfig:"debug"`
	Service        Service     `ocisConfig:"service"`
	Tracing        Tracing     `ocisConfig:"tracing"`
	Ldap           Ldap        `ocisConfig:"ldap"`
	Ldaps          Ldaps       `ocisConfig:"ldaps"`
	Backend        Backend     `ocisConfig:"backend"`
	Fallback       Backend     `ocisConfig:"fallback"`
	Version        string      `ocisConfig:"version"`
	RoleBundleUUID string      `ocisConfig:"role_bundle_uuid"`

	Context    context.Context
	Supervised bool
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}

func DefaultConfig() *Config {
	return &Config{
		Debug: Debug{
			Addr: "127.0.0.1:9129",
		},
		Tracing: Tracing{
			Type:    "jaeger",
			Service: "glauth",
		},
		Service: Service{
			Name: "glauth",
		},
		Ldap: Ldap{
			Enabled:   true,
			Addr:      "127.0.0.1:9125",
			Namespace: "com.owncloud.ldap",
		},
		Ldaps: Ldaps{
			Enabled:   true,
			Addr:      "127.0.0.1:9126",
			Namespace: "com.owncloud.ldaps",
			Cert:      path.Join(defaults.BaseDataPath(), "ldap", "ldap.crt"),
			Key:       path.Join(defaults.BaseDataPath(), "ldap", "ldap.key"),
		},
		Backend: Backend{
			Datastore:   "accounts",
			BaseDN:      "dc=ocis,dc=test",
			Insecure:    false,
			NameFormat:  "cn",
			GroupFormat: "ou",
			Servers:     nil,
			SSHKeyAttr:  "sshPublicKey",
			UseGraphAPI: true,
		},
		Fallback: Backend{
			Datastore:   "",
			BaseDN:      "dc=ocis,dc=test",
			Insecure:    false,
			NameFormat:  "cn",
			GroupFormat: "ou",
			Servers:     nil,
			SSHKeyAttr:  "sshPublicKey",
			UseGraphAPI: true,
		},
		RoleBundleUUID: "71881883-1768-46bd-a24d-a356a2afdf7f", // BundleUUIDRoleAdmin
	}
}
