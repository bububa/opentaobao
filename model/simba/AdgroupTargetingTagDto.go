package simba

// AdgroupTargetingTagDto 结构体
type AdgroupTargetingTagDto struct {
	// 人群包出价方式0:出价;1:溢价
	PriceMode int64 `json:"price_mode,omitempty" xml:"price_mode,omitempty"`
	// 用户溢价比例,溢价20%,接口返回120
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 是否溢价1:不溢价,0:溢价
	IsDefaultPrice int64 `json:"is_default_price,omitempty" xml:"is_default_price,omitempty"`
	// 人群信息
	Crowd *CrowdDto `json:"crowd,omitempty" xml:"crowd,omitempty"`
	// 人群Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 人群上下线状态,0:下线;1:上线
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}
