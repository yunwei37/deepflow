package metric

import (
	"fmt"
)

var DB_FIELD_NEW_FLOW = fmt.Sprintf(
	"if(is_new_flow=%d,1,0)", FLOW_LOG_IS_NEW_FLOW,
)
var DB_FIELD_CLOSED_FLOW = fmt.Sprintf(
	"if(closed_type!=%d,1,0)", FLOW_LOG_CLOSE_TYPE_FORCED_REPORT,
)
var DB_FIELD_TCP_ESTABLISH_FAIL = fmt.Sprintf(
	"if(closed_type in %s,1,0)", FLOW_LOG_CLOSE_TYPE_ESTABLISH_EXCEPTION,
)
var DB_FIELD_CLIENT_ESTABLISH_FAIL = fmt.Sprintf(
	"if(closed_type in %s,1,0)", FLOW_LOG_CLOSE_TYPE_ESTABLISH_EXCEPTION_CLIENT,
)
var DB_FIELD_SERVER_ESTABLISH_FAIL = fmt.Sprintf(
	"if(closed_type in %s,1,0)", FLOW_LOG_CLOSE_TYPE_ESTABLISH_EXCEPTION_SERVER,
)
var DB_FIELD_TCP_TRANSFER_FAIL = fmt.Sprintf(
	"if(closed_type in %s,1,0)", FLOW_LOG_CLOSE_TYPE_EXCEPTION,
)
var DB_FIELD_TCP_RST_FAIL = fmt.Sprintf(
	"if(closed_type in %s,1,0)", FLOW_LOG_CLOSE_TYPE_RST,
)
var DB_FIELD_CLIENT_SOURCE_PORT_REUSE = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_CLIENT_PORT_REUSE,
)
var DB_FIELD_CLIENT_SYN_REPEAT = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_CLIENT_SYN_REPEAT,
)
var DB_FIELD_CLIENT_ESTABLISH_OTHER_RST = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_CLIENT_ESTABLISH_RST,
)
var DB_FIELD_SERVER_RESET = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_SERVER_RST,
)
var DB_FIELD_SERVER_SYN_ACK_REPEAT = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_SERVER_SYNACK_REPEAT,
)
var DB_FIELD_SERVER_ESTABLISH_OTHER_RST = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_SERVER_ESTABLISH_RST,
)
var DB_FIELD_CLIENT_RST_FLOW = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_TCP_CLIENT_RST,
)
var DB_FIELD_SERVER_QUEUE_LACK = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_SERVER_QUEUE_LACK,
)
var DB_FIELD_SERVER_RST_FLOW = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_TCP_SERVER_RST,
)
var DB_FIELD_CLIENT_HALF_CLOSE_FLOW = fmt.Sprintf(
	"if(closed_type=%d,1,0", FLOW_LOG_CLOSE_TYPE_CLIENT_HALF_CLOSE,
)
var DB_FIELD_SERVER_HALF_CLOSE_FLOW = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_SERVER_HALF_CLOSE,
)
var DB_FIELD_TCP_TIMEOUT = fmt.Sprintf(
	"if(closed_type=%d,1,0)", FLOW_LOG_CLOSE_TYPE_TIMEOUT,
)

var L4_FLOW_LOG_METRICS = map[string]*Metric{}

var L4_FLOW_LOG_METRICS_REPLACE = map[string]*Metric{
	"log_count": NewReplaceMetric("1", ""),
	"byte":      NewReplaceMetric("byte_tx+byte_rx", ""),
	"packet":    NewReplaceMetric("packet_tx+packet_rx", ""),
	"l3_byte":   NewReplaceMetric("l3_byte_tx+l3_byte_rx", ""),
	"l4_byte":   NewReplaceMetric("l4_byte_tx+l4_byte_rx", ""),
	"bbp":       NewReplaceMetric("(byte_tx+byte_rx)/(packet_tx+packet_rx)", ""),
	"bbp_tx":    NewReplaceMetric("byte_tx/packet_tx", ""),
	"bbp_rx":    NewReplaceMetric("byte_rx/packet_rx", ""),

	"retrans":           NewReplaceMetric("retrans_tx+retrans_rx", ""),
	"zero_win":          NewReplaceMetric("zero_win_tx+zero_win_rx", ""),
	"retrans_ratio":     NewReplaceMetric("(retrans_tx+retrans_rx)/(packet_tx+packet_rx)", ""),
	"retrans_tx_ratio":  NewReplaceMetric("retrans_tx/packet_tx", ""),
	"retrans_rx_ratio":  NewReplaceMetric("retrans_rx/packet_rx", ""),
	"zero_win_ratio":    NewReplaceMetric("(zero_win_tx+zero_win_rx)/(packet_tx+packet_rx)", ""),
	"zero_win_tx_ratio": NewReplaceMetric("zero_win_tx/packet_tx", ""),
	"zero_win_rx_ratio": NewReplaceMetric("zero_win_rx/packet_rx", ""),

	"new_flow":    NewReplaceMetric(DB_FIELD_NEW_FLOW, ""),
	"closed_flow": NewReplaceMetric(DB_FIELD_CLOSED_FLOW, ""),

	"tcp_establish_fail":          NewReplaceMetric(DB_FIELD_TCP_ESTABLISH_FAIL, ""),
	"client_establish_fail":       NewReplaceMetric(DB_FIELD_CLIENT_ESTABLISH_FAIL, ""),
	"server_establish_fail":       NewReplaceMetric(DB_FIELD_SERVER_ESTABLISH_FAIL, ""),
	"tcp_establish_fail_ratio":    NewReplaceMetric(DB_FIELD_TCP_ESTABLISH_FAIL+"/"+DB_FIELD_CLOSED_FLOW, ""),
	"client_establish_fail_ratio": NewReplaceMetric(DB_FIELD_CLIENT_ESTABLISH_FAIL+"/"+DB_FIELD_CLOSED_FLOW, ""),
	"server_establish_fail_ratio": NewReplaceMetric(DB_FIELD_SERVER_ESTABLISH_FAIL+"/"+DB_FIELD_CLOSED_FLOW, ""),

	"tcp_transfer_fail":          NewReplaceMetric(DB_FIELD_TCP_TRANSFER_FAIL, ""),
	"tcp_transfer_fail_ratio":    NewReplaceMetric(DB_FIELD_TCP_TRANSFER_FAIL+"/"+DB_FIELD_CLOSED_FLOW, ""),
	"tcp_rst_fail":               NewReplaceMetric(DB_FIELD_TCP_RST_FAIL, ""),
	"tcp_rst_fail_ratio":         NewReplaceMetric(DB_FIELD_TCP_RST_FAIL+"/"+DB_FIELD_CLOSED_FLOW, ""),
	"client_source_port_reuse":   NewReplaceMetric(DB_FIELD_CLIENT_SOURCE_PORT_REUSE, ""),
	"client_syn_repeat":          NewReplaceMetric(DB_FIELD_CLIENT_SYN_REPEAT, ""),
	"client_establish_other_rst": NewReplaceMetric(DB_FIELD_CLIENT_ESTABLISH_OTHER_RST, ""),
	"server_reset":               NewReplaceMetric(DB_FIELD_SERVER_RESET, ""),
	"server_syn_ack_repeat":      NewReplaceMetric(DB_FIELD_SERVER_SYN_ACK_REPEAT, ""),
	"server_establish_other_rst": NewReplaceMetric(DB_FIELD_SERVER_ESTABLISH_OTHER_RST, ""),
	"client_rst_flow":            NewReplaceMetric(DB_FIELD_CLIENT_RST_FLOW, ""),
	"server_queue_lack":          NewReplaceMetric(DB_FIELD_SERVER_QUEUE_LACK, ""),
	"server_rst_flow":            NewReplaceMetric(DB_FIELD_SERVER_RST_FLOW, ""),
	"client_half_close_flow":     NewReplaceMetric(DB_FIELD_CLIENT_HALF_CLOSE_FLOW, ""),
	"server_half_close_flow":     NewReplaceMetric(DB_FIELD_SERVER_HALF_CLOSE_FLOW, ""),
	"tcp_timeout":                NewReplaceMetric(DB_FIELD_TCP_TIMEOUT, ""),

	"rtt_avg":    NewReplaceMetric("rtt/1", ""),
	"rtt_client": NewReplaceMetric("rtt_client_sum/rtt_client_count", ""),
	"rtt_server": NewReplaceMetric("rtt_server_sum/rtt_server_count", ""),
	"srt":        NewReplaceMetric("srt_sum/srt_count", ""),
	"art":        NewReplaceMetric("art_sum/art_count", ""),
	"rrt":        NewReplaceMetric("rrt_sum/rrt_count", ""),

	"l7_error":              NewReplaceMetric("l7_client_error+l7_server_error", ""),
	"l7_error_ratio":        NewReplaceMetric("l7_request/l7_response", ""),
	"l7_client_error_ratio": NewReplaceMetric("l7_client_error/l7_response", ""),
	"l7_server_error_ratio": NewReplaceMetric("l7_server_error/l7_response", ""),

	"vpc_0":         NewReplaceMetric("[toString(l3_epc_id_0)]", "NOT (l3_epc_id_0 = -2)"),
	"subnet_0":      NewReplaceMetric("[toString(subnet_id_0)]", "NOT (subnet_id_0 = 0)"),
	"ip_0":          NewReplaceMetric("[toString(ip4_0), toString(subnet_id_0), toString(is_ipv4), toString(ip6_0)]", "NOT (((is_ipv4 = 1) OR (ip6_0 = toIPv6('::'))) AND ((is_ipv4 = 0) OR (ip4_0 = toIPv4('0.0.0.0'))))"),
	"pod_cluster_0": NewReplaceMetric("[toString(pod_cluster_id_0)]", "NOT (pod_cluster_id_0 = 0)"),
	"pod_node_0":    NewReplaceMetric("[toString(pod_node_id_0)]", "NOT (pod_node_id_0 = 0)"),
	"pod_ns_0":      NewReplaceMetric("[toString(pod_ns_id_0)]", "NOT (pod_ns_id_0 = 0)"),
	"pod_group_0":   NewReplaceMetric("[toString(pod_group_id_0)]", "NOT (pod_group_id_0 = 0)"),
	"pod_0":         NewReplaceMetric("[toString(pod_id_0)]", "NOT (pod_id_0 = 0)"),
	"host_0":        NewReplaceMetric("[toString(host_id_0)]", "NOT (host_id_0 = 0)"),
	"vm_0":          NewReplaceMetric("[toString(l3_device_id_0), toString(l3_device_type_0)]", "(NOT (l3_device_id_0 = 0)) AND (l3_device_type_0 = 1)"),
	"region_0":      NewReplaceMetric("[toString(region_id_0)]", "NOT (region_id_0 = 0)"),
	"az_0":          NewReplaceMetric("[toString(az_id_0)]", "NOT (az_id_0 = 0)"),
	"vpc_1":         NewReplaceMetric("[toString(l3_epc_id_1)]", "NOT (l3_epc_id_1 = -2)"),
	"subnet_1":      NewReplaceMetric("[toString(subnet_id_1)]", "NOT (subnet_id_1 = 0)"),
	"ip_1":          NewReplaceMetric("[toString(ip4_1), toString(subnet_id_1), toString(is_ipv4), toString(ip6_1)]", "NOT (((is_ipv4 = 1) OR (ip6_1 = toIPv6('::'))) AND ((is_ipv4 = 0) OR (ip4_1 = toIPv4('0.0.0.0'))))"),
	"pod_cluster_1": NewReplaceMetric("[toString(pod_cluster_id_1)]", "NOT (pod_cluster_id_1 = 0)"),
	"pod_node_1":    NewReplaceMetric("[toString(pod_node_id_1)]", "NOT (pod_node_id_1 = 0)"),
	"pod_ns_1":      NewReplaceMetric("[toString(pod_ns_id_1)]", "NOT (pod_ns_id_1 = 0)"),
	"pod_group_1":   NewReplaceMetric("[toString(pod_group_id_1)]", "NOT (pod_group_id_1 = 0)"),
	"pod_1":         NewReplaceMetric("[toString(pod_id_1)]", "NOT (pod_id_1 = 0)"),
	"host_1":        NewReplaceMetric("[toString(host_id_1)]", "NOT (host_id_1 = 0)"),
	"vm_1":          NewReplaceMetric("[toString(l3_device_id_1), toString(l3_device_type_1)]", "(NOT (l3_device_id_1 = 0)) AND (l3_device_type_1 = 1)"),
	"region_1":      NewReplaceMetric("[toString(region_id_1)]", "NOT (region_id_1 = 0)"),
	"az_1":          NewReplaceMetric("[toString(az_id_1)]", "NOT (az_id_1 = 0)"),
}

func GetL4FlowLogMetrics() map[string]*Metric {
	// TODO: 特殊指标量修改
	return L4_FLOW_LOG_METRICS
}
