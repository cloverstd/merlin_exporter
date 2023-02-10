package merlin_client

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var uptimePattern = regexp.MustCompile("\\((\\d+) secs since boot\\)")

type DeviceMap struct {
	Sys string `xml:"sys"`
}

func (mc *MerlinClient) Uptime(ctx context.Context) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mc.renderURL("ajax_status.xml"), nil)
	if err != nil {
		return 0, err
	}
	resp, err := mc.do(ctx, req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return parseUptime(resp.Body)
}

func parseUptime(r io.Reader) (int64, error) {
	dec := xml.NewDecoder(r)
	var dm DeviceMap
	err := dec.Decode(&dm)
	if err != nil {
		return 0, err
	}
	submatch := uptimePattern.FindStringSubmatch(dm.Sys)
	if len(submatch) != 2 {
		return 0, fmt.Errorf("invalid uptime info, [%s]", dm.Sys)
	}
	return strconv.ParseInt(submatch[1], 10, 64)
}
