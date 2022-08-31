package product

// ClientInfoDto 结构体
type ClientInfoDto struct {
	// 平台ID,2:&#34;安卓&#34;,3:&#34;苹果&#34;,4,:&#34;越狱&#34;,1:&#34;其他&#34;,5:&#34;PC&#34;,6:&#34;XBOX - 主机&#34;
	PlatformId int64 `json:"platform_id,omitempty" xml:"platform_id,omitempty"`
	// 客户端ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
