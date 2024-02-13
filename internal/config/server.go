package config

type server struct {
	port uint
}

func (s server) GetPort() uint {
	return s.port
}
