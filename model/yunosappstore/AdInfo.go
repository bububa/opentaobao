package yunosappstore

// AdInfo 结构体
type AdInfo struct {
	// 广告跟踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 简介
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 版本号
	VerCode int64 `json:"ver_code,omitempty" xml:"ver_code,omitempty"`
	// 版本名
	VerName string `json:"ver_name,omitempty" xml:"ver_name,omitempty"`
	// 结算类型
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 包名
	Pkg string `json:"pkg,omitempty" xml:"pkg,omitempty"`
	// 包文件md5
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 回调地址
	Callbacks *Callbacks `json:"callbacks,omitempty" xml:"callbacks,omitempty"`
	// 创意列表
	Creatives []string `json:"creatives,omitempty" xml:"creatives>string,omitempty"`
	// deeplink
	Deeplink string `json:"deeplink,omitempty" xml:"deeplink,omitempty"`
	// 下地地址
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
}
