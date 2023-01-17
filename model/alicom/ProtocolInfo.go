package alicom

// ProtocolInfo 结构体
type ProtocolInfo struct {
	// 协议列表
	ProtocolList []ProtocolList `json:"protocol_list,omitempty" xml:"protocol_list>protocol_list,omitempty"`
	// 协议查询流水id
	ProtocolSequenceId string `json:"protocol_sequence_id,omitempty" xml:"protocol_sequence_id,omitempty"`
}
