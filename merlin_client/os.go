package merlin_client

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

type MemoryInfo struct {
	Total int `xml:"total"`
	Free  int `xml:"free"`
	Used  int `xml:"used"`
}

type CPUInfo struct {
	Total int `xml:"total"`
	Usage int `xml:"usage"`
}

type OSInfo struct {
	Memory *MemoryInfo `xml:"mem_info"`
	CPU    []CPUInfo   `xml:"cpu_info>cpu"`
}

func parseOSInfo(r io.Reader) (*OSInfo, error) {
	dec := xml.NewDecoder(r)
	var info OSInfo
	err := dec.Decode(&info)
	return &info, err
}

func (mc *MerlinClient) OSInfo(ctx context.Context) (*OSInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mc.renderURL("cpu_ram_status.xml"), nil)
	if err != nil {
		return nil, err
	}
	body, err := mc.do(ctx, req)
	if err != nil {
		return nil, err
	}
	return parseOSInfo(bytes.NewReader(body))
}
