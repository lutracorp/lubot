package database

type Config struct {
	Type string `hcl:"type,label"`
	DSN  string `hcl:"dsn"`
}
