package trade

// CommonKeyValue 结构体
type CommonKeyValue struct {
	// 扩展key。例如：.cpCode（物流品牌）。详以接入文档中描述的场景对接
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 扩展value。例如：传运单号。详以接入文档中描述的场景对接
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
