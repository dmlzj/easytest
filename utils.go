package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	arrayJSONPrefix  byte = '['
	objectJSONPrefix byte = '{'
)

func commandsFromJSON(msgBytes []byte) ([]Command, error) {
	var cmds []Command
	firstByte := msgBytes[0]
	switch firstByte {
	case objectJSONPrefix:
		// single command request
		var cmd Command
		err := json.Unmarshal(msgBytes, &cmd)
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

func checkvalue(k []string, v interface{}, m map[string]interface{}) (interface{}, error) {
	if len(k) == 1 {
		log.Println("checkvalue:len==1:", k, v)
		if mv, ok := m[k[0]]; ok {
			if fmt.Sprint(v) == fmt.Sprint(mv) {
				return v, nil
			} else {
				return nil, fmt.Errorf("Key:%s Value:%v != %v", k[0], v, mv)
			}
		} else {
			return nil, fmt.Errorf("No Key:%s", k[0])
		}
	}

	if mv, ok := m[k[0]]; ok {
		if nm, ok := mv.(map[string]interface{}); ok {
			return checkvalue(k[1:], v, nm)
		}
	}
	return nil, fmt.Errorf("No Key:%s", k[0])
}
