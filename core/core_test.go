package core_test

import (
	core "github.com/misostack/ezgo/core"
	"testing"
	"fmt"
)

func TestParseWebServerConfig(t *testing.T) {
	cfg := core.WebServerConfig{}
	core.ParseWebServerConfig(&cfg)
	fmt.Printf("%v\n", cfg)	
}

// func TestConfigStructToMap(t *testing.T) {
// 	// test with Config struct
// 	cfg := core.Config{}
// 	cfgm := make(map[string]interface{})
// 	cfg.StructToMap(&cfgm)
// 	expectedMap := map[string]string{
// 		"Port": "",
// 		"Address": "",
// 	}
// 	is_valid := true
// 	for k, _ := range expectedMap {
// 		_, ok := cfgm[k]
// 		if !ok {
// 			is_valid = false
// 			break
// 		}
// 	}
// 	if !is_valid {
// 		t.Errorf("Config.StructToMap(m) failed!")
// 	}
// }

// func TestSetStructField(t *testing.T) {
// 	s := core.Config{
// 		Address: "127.0.0.1",
// 		Port: "8081",
// 	}

// 	// reflect.ValueOf(&s).Elem().FieldByName("Address").Set(reflect.ValueOf("something"))
// 	core.SetStructField(s, "Address", "10.84.4.64")
// 	fmt.Printf("%v", s)
// }