package product

// ClientInfoDto 结构体
type ClientInfoDto struct {
	// 平台ID,2:"安卓",3:"苹果",4,:"越狱",1:"其他",5:"PC",6:"XBOX - 主机"
	PlatformId int64 `json:"platform_id,omitempty" xml:"platform_id,omitempty"`
	// 客户端ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
