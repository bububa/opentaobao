package wlb

// WlbPartnerContact 结构体
type WlbPartnerContact struct {
	// 仓库联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}
