package main

import (
	"encoding/json"
)

var (
	arrayJSONPrefix  byte = '['
	objectJSONPrefix byte = '{'
)

func commandsFromJSON(msgBytes []byte) ([]*Command, error) {
	cmds := []*Command{}
	firstByte := msgBytes[0]
	switch firstByte {
	case objectJSONPrefix:
		// single command request
		cmd := &Command{}
		err := json.Unmarshal(msgBytes, cmd)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	case arrayJSONPrefix:
		// array of commands received
		err := json.Unmarshal(msgBytes, &cmds)
		if err != nil {
			return nil, err
		}
	}
	return cmds, nil
}

// func checkvalue(k []string, v interface{}, m map[string]interface{}) interface{} {
// 	kk := strings.Split(k[0], ":")
// 	if len(k) == 1 {
// 		//log.Println("checkvalue:len==1:", k, v, m)
// 		if mv, ok := m[kk[0]]; ok {
// 			//	log.Println("checkvalue:mv:", kk[0], mv, m, reflect.TypeOf(mv))
// 			rt := reflect.TypeOf(mv)
// 			switch rt.Kind() {
// 			case reflect.Slice, reflect.Array:
// 				index, err := strconv.Atoi(kk[1])
// 				if err != nil {
// 					panic(err)
// 				}
// 				mvr := reflect.ValueOf(mv).Index(index)
// 				if fmt.Sprint(v) == fmt.Sprint(mvr) {
// 					return mv
// 				}
// 				panic(fmt.Errorf("Key:%s Value:%v[%v] != %v[%v]", k[0], v, reflect.TypeOf(v).Name(), mvr, mvr.Type().Name()))
// 			default:
// 				if reflect.TypeOf(v).Name() == reflect.TypeOf(mv).Name() && fmt.Sprint(v) == fmt.Sprint(mv) {
// 					return mv
// 				}
// 				log.Println("checkvalue:", reflect.TypeOf(v).Name(), reflect.TypeOf(mv).Name(), v, mv)
// 				panic(fmt.Errorf("Key:%s Value:%v != %v", k[0], v, mv))
// 			}
// 		}
// 		panic(fmt.Errorf("No Key:%s", k[0]))
// 	}
// 	if mv, ok := m[kk[0]]; ok {
// 		switch t := mv.(type) {
// 		case map[string]interface{}:
// 			return checkvalue(k[1:], v, t)
// 		case []map[string]interface{}:
// 			index, err := strconv.Atoi(kk[1])
// 			if err != nil {
// 				panic(err)
// 			}
// 			return checkvalue(k[1:], v, t[index])
// 		default:
// 			panic(fmt.Errorf("Key %s don't has children.%+v", kk, t))
// 		}
// 	}
// 	panic(fmt.Errorf("No Key:%s", k[0]))
// }
