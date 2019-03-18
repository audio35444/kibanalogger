package kibanalogger

import (
	"fmt"
	"reflect"
)

/*
KibanaLogs use in String method
*/
func KibanaLogs(s interface{}) (logs string) {
	loopToStruct(s, "kibana", func(tag, name, value string) {
		logs += "[" + tag + ": " + value + "]"
	})
	return
}

func loopToStruct(s interface{}, tagSearch string, callback func(tag, name, value string)) {
	elem := reflect.TypeOf(s).Elem()
	val := reflect.ValueOf(s).Elem()
	limit := elem.NumField()
	for i := 0; i < limit; i++ {
		field := elem.Field(i)
		tag := field.Tag.Get(tagSearch)
		if tag != "" {
			name := field.Name
			value := val.Field(i)
			callback(tag, name, fmt.Sprint(value)) //TODO: hacer sin usar fmt.Sprint
		}
	}
}
