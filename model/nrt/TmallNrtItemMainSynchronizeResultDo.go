package nrt

// TmallnrtitemmainsynchronizeResultDo 结构体
type TmallnrtitemmainsynchronizeResultDo struct {
	// 返回值
	Data *NrtItemSyncResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
