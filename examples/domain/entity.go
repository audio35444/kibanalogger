package domain

import "github.com/audio35444/kibanalogger"

type Entity struct {
	Id          uint64
	Name        string `kibana:"entity_name"`
	Description string `kibana:"description"`
}

func (e Entity) String() string {
	return kibanalogger.KibanaLogs(&e)
}
