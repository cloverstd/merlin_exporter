package merlin_client

import (
	"bytes"
	"testing"
)

func TestMerlinClient_parseOSInfo(t *testing.T) {
	info, err := parseOSInfo(bytes.NewBufferString(`<?xml version="1.0" ?>
<info>
<cpu_info>
<cpu>
<total>1540956113</total>
<usage>10121455</usage>
</cpu>
<cpu>
<total>1540956074</total>
<usage>17635635</usage>
</cpu>
</cpu_info>

<mem_info>
<total>255716</total>
<free>184248</free>
<used>71468</used>
</mem_info>

</info>`))

	if err != nil {
		t.Fatalf("parse failed, %v", err)
	}
	if info.Memory.Total != 255716 && info.Memory.Free != 184248 && info.Memory.Used != 71468 {
		t.Fatalf("parse memeory failed")
	}

	if info.CPU[1].Total != 1540956074 && info.CPU[1].Usage != 17635635 {
		t.Fatalf("parse cpu failed")
	}

	if info.CPU[0].Total != 1540956113 && info.CPU[0].Usage != 10121455 {
		t.Fatalf("parse cpu failed")
	}
}
