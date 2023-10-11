package idle

// IdleAdvMaterialUploadTopResult 结构体
type IdleAdvMaterialUploadTopResult struct {
	// 失败的素材的名称，和参数传入的title对应
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 当前错误原因
	Result *IdleAdvBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
