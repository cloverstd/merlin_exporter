package merlin_client

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (mc *MerlinClient) Temperature(ctx context.Context) (map[string]float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mc.renderURL("ajax_coretmp.asp"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := mc.do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parseTemperature(resp.Body)
}

func parseTemperature(body io.Reader) (map[string]float64, error) {
	scanner := bufio.NewScanner(body)
	result := map[string]float64{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "disabled") {
			continue
		}
		fields := strings.Split(text, "=")
		if len(fields) != 2 {
			continue
		}
		float, err := strconv.ParseFloat(strings.NewReplacer(
			"\"", "",
			";", "",
			"&deg", "",
			"C", "",
			" ", "",
		).Replace(fields[1]), 64)
		if err != nil {
			return nil, err
		}
		result[strings.Trim(strings.TrimSpace(fields[0]), "_raw")] = float
	}
	return result, nil
}
