package merlin_client

import (
	"bytes"
	"testing"
)

func TestMerlinClient_Temperature(t *testing.T) {
	temperature, err := parseTemperature(bytes.NewBufferString(`curr_coreTmp_2_raw = "44&deg;C";
curr_coreTmp_2 = (curr_coreTmp_2_raw.indexOf("disabled") > 0 ? 0 : curr_coreTmp_2_raw.replace("&deg;C", ""));
curr_coreTmp_5_raw = "48&deg;C";
curr_coreTmp_5 = (curr_coreTmp_5_raw.indexOf("disabled") > 0 ? 0 : curr_coreTmp_5_raw.replace("&deg;C", ""));
curr_coreTmp_cpu = "61";

`))
	if err != nil {
		t.Fatalf("error should be nil, but is %v", err)
	}
	if temperature["curr_coreTmp_2"] != 44 {
		t.Fatalf("parse curr_coreTmp_2 failed")
	}
	if temperature["curr_coreTmp_5"] != 48 {
		t.Fatalf("parse curr_coreTmp_5 failed")
	}
	if temperature["curr_coreTmp_cpu"] != 61 {
		t.Fatalf("parse curr_coreTmp_cpu failed")
	}
}
