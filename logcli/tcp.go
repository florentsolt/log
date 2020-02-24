package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

type Hosts struct {
	hosts []*Host
}

var re = regexp.MustCompile(`([\w\d\.\-_]+)@([^\s:]+):(\d+)`)

func ParseHosts(strings []string) (*Hosts, error) {
	result := &Hosts{}
	maxLength := 0
	for _, s := range strings {
		matches := re.FindStringSubmatch(s)
		if len(matches) != 4 {
			return nil, fmt.Errorf(`Malformed hostname@ip:port "%s" %#v`, s, matches)
		}
		if len(matches[1]) > maxLength {
			maxLength = len(matches[1])
		}
		result.hosts = append(result.hosts, &Host{name: matches[1], nameWithPaddingAndSpace: []byte(matches[1]), ip: matches[2], port: matches[3]})
	}

	l := strconv.Itoa(maxLength)
	for _, h := range result.hosts {
		h.nameWithPaddingAndSpace = []byte(fmt.Sprintf("%"+l+"s ", h.name))
	}
	return result, nil
}

func (hh *Hosts) IsEmpty() bool {
	return len(hh.hosts) == 0
}

func (hh *Hosts) Connect(timeout time.Duration) (Channel, error) {
	channel := make(Channel)
	for _, host := range hh.hosts {
		err := host.Connect(timeout)
		if err != nil {
			return nil, fmt.Errorf("Unable to connect to %s", host.String())
		}
	}
	// wait for all connection to start reading
	for _, host := range hh.hosts {
		go host.Read(channel)
	}
	return channel, nil
}

type Message struct {
	HostnameWithPaddingAndSpace []byte
	Line                        []byte
}

type Channel chan Message

type Host struct {
	name                    string
	nameWithPaddingAndSpace []byte
	ip                      string
	port                    string
	conn                    net.Conn
}

func (h *Host) String() string {
	return h.name + "@" + h.ip + ":" + h.port
}

func (h *Host) Connect(timeout time.Duration) (err error) {
	h.conn, err = net.DialTimeout("tcp4", h.ip+":"+h.port, timeout)
	return
}

func (h *Host) Read(channel Channel) {
	for {
		scanner := bufio.NewScanner(h.conn)
		for scanner.Scan() {
			channel <- Message{h.nameWithPaddingAndSpace, scanner.Bytes()}
		}
	}
}
