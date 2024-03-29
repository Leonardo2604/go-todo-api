package config

import "fmt"

type database struct {
	host string
	user string
	pass string
	name string
	port uint
}

func (d database) GetHost() string {
	return d.host
}

func (d database) GetUser() string {
	return d.user
}

func (d database) GetPass() string {
	return d.pass
}

func (d database) GetName() string {
	return d.name
}

func (d database) GetPort() uint {
	return d.port
}

func (d database) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", d.host, d.user, d.pass, d.name, d.port)
}
