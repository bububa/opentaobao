package wdk

// PropField 结构体
type PropField struct {
	// 渠道属性value，取值见key定义
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 渠道属性key, 取值为"ONE_HOUR_STATUS" 代表小时达，value=0表示不可售, value=1表示可售；"WAVE_ARRIVE_STATUS" 代表波次达，value=0表示不可售,value=1表示可售
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
