//
//  Daemon for IVPN Client Desktop
//  https://github.com/ivpn/desktop-app-daemon
//
//  Created by Stelnykovych Alexandr.
//  Copyright (c) 2020 Privatus Limited.
//
//  This file is part of the Daemon for IVPN Client Desktop.
//
//  The Daemon for IVPN Client Desktop is free software: you can redistribute it and/or
//  modify it under the terms of the GNU General Public License as published by the Free
//  Software Foundation, either version 3 of the License, or (at your option) any later version.
//
//  The Daemon for IVPN Client Desktop is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
//  or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
//  details.
//
//  You should have received a copy of the GNU General Public License
//  along with the Daemon for IVPN Client Desktop. If not, see <https://www.gnu.org/licenses/>.
//

package types

// WireGuardServerHostInfo contains info about WG server host
type WireGuardServerHostInfo struct {
	Hostname  string `json:"hostname"`
	Host      string `json:"host"`
	PublicKey string `json:"public_key"`
	LocalIP   string `json:"local_ip"`
}

// WireGuardServerInfo contains all info about WG server
type WireGuardServerInfo struct {
	Gateway     string `json:"gateway"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	City        string `json:"city"`

	Hosts []WireGuardServerHostInfo `json:"hosts"`
}

// OpenvpnServerInfo contains all info about OpenVPN server
type OpenvpnServerInfo struct {
	Gateway     string   `json:"gateway"`
	CountryCode string   `json:"country_code"`
	Country     string   `json:"country"`
	City        string   `json:"city"`
	IPAddresses []string `json:"ip_addresses"`
}

// DNSInfo contains info about DNS server
type DNSInfo struct {
	IP         string `json:"ip"`
	MultihopIP string `json:"multihop-ip"`
}

// AntitrackerInfo all info about antitracker DNSs
type AntitrackerInfo struct {
	Default  DNSInfo `json:"default"`
	Hardcore DNSInfo `json:"hardcore"`
}

// InfoAPI contains API IP adresses
type InfoAPI struct {
	IPAddresses []string `json:"ips"`
}

// ConfigInfo contains different configuration info (Antitracker, API ...)
type ConfigInfo struct {
	Antitracker AntitrackerInfo `json:"antitracker"`
	API         InfoAPI         `json:"api"`
}

// ServersInfoResponse all info from servers.json
type ServersInfoResponse struct {
	WireguardServers []WireGuardServerInfo `json:"wireguard"`
	OpenvpnServers   []OpenvpnServerInfo   `json:"openvpn"`
	Config           ConfigInfo            `json:"config"`
}
