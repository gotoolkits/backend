package backend

import (
	"fmt"
)

type backend struct {
	ver   string
	gover string
}

var (
	Backend = &backend{
		ver:   "0.1",
		gover: "1.5+",
	}
)

func (s *backend) Description() string {
	return fmt.Sprintf("backend verion %s, build on go %s", s.ver, s.gover)
}
