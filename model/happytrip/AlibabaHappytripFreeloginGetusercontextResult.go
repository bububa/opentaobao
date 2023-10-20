package happytrip

// AlibabahappytripfreelogingetusercontextResult 结构体
type AlibabahappytripfreelogingetusercontextResult struct {
	// 错误消息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 返回的结果值
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}
