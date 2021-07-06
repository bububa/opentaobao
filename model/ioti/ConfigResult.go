package ioti

// ConfigResult 结构体
type ConfigResult struct {
	// IPC列表
	Ipcs []string `json:"ipcs,omitempty" xml:"ipcs>string,omitempty"`
	// agent标识
	AgentId string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 连接客户端id
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 上报消息topic
	MqttUpTopic string `json:"mqtt_up_topic,omitempty" xml:"mqtt_up_topic,omitempty"`
	// 通道类型
	ChannelType string `json:"channel_type,omitempty" xml:"channel_type,omitempty"`
	// broker代理地址
	BrokerUrl string `json:"broker_url,omitempty" xml:"broker_url,omitempty"`
	// 下发消息topic
	MqttDownTopic string `json:"mqtt_down_topic,omitempty" xml:"mqtt_down_topic,omitempty"`
	// 访问accesskey
	AccessKey string `json:"access_key,omitempty" xml:"access_key,omitempty"`
	// CPU上线
	CpuUpperLimit string `json:"cpu_upper_limit,omitempty" xml:"cpu_upper_limit,omitempty"`
	// 是否保活
	KeepAlive int64 `json:"keep_alive,omitempty" xml:"keep_alive,omitempty"`
	// 最大推流数量
	MaxStreamNumber int64 `json:"max_stream_number,omitempty" xml:"max_stream_number,omitempty"`
	// 上报时间间隔
	ReportInterval int64 `json:"report_interval,omitempty" xml:"report_interval,omitempty"`
}
