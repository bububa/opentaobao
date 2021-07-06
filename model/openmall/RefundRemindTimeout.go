package openmall

// RefundRemindTimeout 结构体
type RefundRemindTimeout struct {
	// 超时时间。格式:yyyy-MM-dd HH:mm:ss
	Timeout string `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 是否存在超时。可选值:true(是),false(否)
	ExistTimeout bool `json:"exist_timeout,omitempty" xml:"exist_timeout,omitempty"`
}
