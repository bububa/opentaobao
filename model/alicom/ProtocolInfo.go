package alicom

import (
	"sync"
)

// ProtocolInfo 结构体
type ProtocolInfo struct {
	// 协议列表
	ProtocolList []ProtocolList `json:"protocol_list,omitempty" xml:"protocol_list>protocol_list,omitempty"`
	// 协议查询流水id
	ProtocolSequenceId string `json:"protocol_sequence_id,omitempty" xml:"protocol_sequence_id,omitempty"`
}

var poolProtocolInfo = sync.Pool{
	New: func() any {
		return new(ProtocolInfo)
	},
}

// GetProtocolInfo() 从对象池中获取ProtocolInfo
func GetProtocolInfo() *ProtocolInfo {
	return poolProtocolInfo.Get().(*ProtocolInfo)
}

// ReleaseProtocolInfo 释放ProtocolInfo
func ReleaseProtocolInfo(v *ProtocolInfo) {
	v.ProtocolList = v.ProtocolList[:0]
	v.ProtocolSequenceId = ""
	poolProtocolInfo.Put(v)
}
