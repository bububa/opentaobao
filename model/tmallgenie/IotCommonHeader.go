package tmallgenie

// IotCommonHeader 
type IotCommonHeader struct {
    // 标准控制协议中的namespace
    Namespace   string `json:"namespace,omitempty" xml:"namespace,omitempty"`
    // 标准控制协议中的name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 标准控制协议中的messageId
    MessageId   string `json:"message_id,omitempty" xml:"message_id,omitempty"`
    // 标准控制协议中的payLoadVersion
    PayLoadVersion   int64 `json:"pay_load_version,omitempty" xml:"pay_load_version,omitempty"`
}
