package tmc

// TmcQueueInfo 结构体
type TmcQueueInfo struct {
	// 当前队列当天读取量
	GetTotal int64 `json:"get_total,omitempty" xml:"get_total,omitempty"`
	// 当前队列当天写入量
	PutToal int64 `json:"put_toal,omitempty" xml:"put_toal,omitempty"`
	// TMC组名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 消息队列Broker名称
	BrokerName string `json:"broker_name,omitempty" xml:"broker_name,omitempty"`
}
