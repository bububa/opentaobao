package security

// JaqDispatchParam 
type JaqDispatchParam struct {

    // 事件ID
    EventId   string `json:"event_id,omitempty"`

    // 协议版本号
    ProtocolVersion   string `json:"protocol_version,omitempty"`

    // 下发的软token密文
    Rtken   string `json:"rtken,omitempty"`

    // 下发的软token索引
    RtkenIndex   string `json:"rtken_index,omitempty"`

}
