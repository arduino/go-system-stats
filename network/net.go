//
//  This file is part of go-system-stats library
//
//  Copyright (C) 2017  Arduino AG (http://www.arduino.cc/)
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package net

import (
	"github.com/arduino/gonetworkmanager"
)

// Stats contains data about the memory avaible/used
type Stats struct {
	Devices []gonetworkmanager.Device
	Status  string
}

type IpProxyConfig struct {
	Manual bool
	Config gonetworkmanager.IpProxyConfig
}

func GetNetworkStats() (*Stats, error) {
	res := &Stats{}
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return nil, err
	}
	res.Devices = nm.GetDevices()
	res.Status = nm.GetState().String()
	return res, nil
}

func AddWirelessConnection(ssid, password string) error {
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return err
	}
	nm.AddWirelessConnection(ssid, password)
	return nil
}

func AddWiredConnection(config IpProxyConfig) error {
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return err
	}
	nm.AddWiredConnection(config.Manual, config.Config)
	return nil
}
