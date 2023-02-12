package merlin_client

import (
	"bytes"
	"testing"
)

func TestMerlinClient_parseUptime(t *testing.T) {
	uptime, err := parseUptime(bytes.NewBufferString(`<?xml version="1.0" ?>
<devicemap>  <wan>0</wan>
<wan>0</wan>
<wan>0</wan>

  <wan>monoClient=</wan>
  <wan>wlc_state=0</wan>
  <wan>wlc_sbstate=0</wan>
  <wan>psta:wlc_state=0;wlc_state_auth=0;</wan>
  <wan>wifi_hw_switch=1</wan>
  <wan>ddnsRet=</wan>
  <wan>ddnsUpdate=0</wan>
  <wan>wan_line_state=</wan>
  <wan>wlan0_radio_flag=1</wan>
  <wan>wlan1_radio_flag=1</wan>
  <wan>wlan2_radio_flag=</wan>
  <wan>data_rate_info_2g=72 Mbps</wan>  
  <wan>data_rate_info_5g=468 Mbps</wan>  
  <wan>data_rate_info_5g_2=0 Mbps</wan>
  <wan>wan_diag_state=</wan>
  <wan>active_wan_unit=0</wan>
  <wan>wan0_enable=1</wan>
  <wan>wan1_enable=1</wan>
  <wan>wan0_realip_state=0</wan>
  <wan>wan1_realip_state=0</wan>
  <wan>wan0_ipaddr=0.0.0.0</wan>
  <wan>wan1_ipaddr=0.0.0.0</wan>
  <wan>wan0_realip_ip=</wan>
  <wan>wan1_realip_ip=</wan>
  <vpn>vpnc_proto=disable</vpn>
  <vpn>vpnc_state_t=0</vpn>
  <vpn>vpnc_sbstate_t=0</vpn>
  <vpn>vpn_client1_state=0</vpn>
  <vpn>vpn_client2_state=0</vpn>
  <vpn>vpn_client3_state=0</vpn>
  <vpn>vpn_client4_state=0</vpn>
  <vpn>vpn_client5_state=0</vpn>
  <vpn>vpnd_state=0</vpn>
  <vpn>vpn_client1_errno=0</vpn>
  <vpn>vpn_client2_errno=0</vpn>
  <vpn>vpn_client3_errno=0</vpn>
  <vpn>vpn_client4_errno=0</vpn>
  <vpn>vpn_client5_errno=0</vpn>
  <vpn>vpn_server1_state=0</vpn>
  <vpn>vpn_server2_state=0</vpn>
  <sys>uptimeStr=Thu, 09 Feb 2023 22:05:18 +0800(15411660 secs since boot)</sys>
  <qtn>qtn_state=</qtn>
  <usb>'[]'</usb>
  <usb>modem_enable=1</usb>
  <first_wan>0</first_wan>
<first_wan>0</first_wan>
<first_wan>0</first_wan>
<second_wan>0</second_wan>
<second_wan>0</second_wan>
<second_wan>0</second_wan>

  <sim>sim_state=</sim>
  <sim>sim_signal=</sim>
  <sim>sim_operation=</sim>
  <sim>sim_isp=</sim>
  <sim>roaming=0</sim>
  <sim>roaming_imsi=</sim>
  <sim>sim_imsi=</sim>
  <sim>g3err_pin=</sim>
  <sim>pin_remaining_count=</sim>
  <sim>modem_act_provider=</sim>
  <sim>rx_bytes=</sim>
  <sim>tx_bytes=</sim>
  <sim>modem_sim_order=</sim>
  <dhcp>dnsqmode=</dhcp>
  <wan>wlc0_state=</wan>
  <wan>wlc1_state=</wan>
  <wan>rssi_2g=</wan>
  <wan>rssi_5g=</wan>
  <wan>rssi_5g_2=</wan>
</devicemap>
`))
	if err != nil {
		t.Fatalf("parse uptime failed, %v", err)
	}

	if uptime != 1675951517 {
		t.Fatalf("parse uptime value failed, %v", uptime)
	}
}
