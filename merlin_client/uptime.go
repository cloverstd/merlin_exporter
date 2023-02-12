package merlin_client

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DeviceMap struct {
	Sys string `xml:"sys"`
}

func (mc *MerlinClient) Uptime(ctx context.Context) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mc.renderURL("ajax_status.xml"), nil)
	if err != nil {
		return 0, err
	}
	body, err := mc.do(ctx, req)
	if err != nil {
		return 0, err
	}
	return parseUptime(bytes.NewReader(body))
}

func parseUptime(r io.Reader) (int64, error) {
	dec := xml.NewDecoder(r)
	var dm DeviceMap
	err := dec.Decode(&dm)
	if err != nil {
		return 0, err
	}
	fields := strings.Split(dm.Sys, "(")
	if len(fields) != 2 {
		return 0, fmt.Errorf("invalid uptime info, [%s]", dm.Sys)
	}

	routerNow, err := time.Parse(time.RFC1123Z, strings.TrimPrefix(fields[0], "uptimeStr="))
	if err != nil {
		return 0, fmt.Errorf("parse uptime now failed, [%s], %v", fields[0], err)
	}

	since, err := strconv.ParseInt(strings.Split(fields[1], " ")[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse uptime since failed, [%s], %v", fields[1], err)
	}

	return routerNow.Add(time.Duration(since) * time.Second * -1).Unix(), nil
}
