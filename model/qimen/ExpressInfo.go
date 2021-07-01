package qimen

// ExpressInfo 结构体
type ExpressInfo struct {
	// 奇门仓储字段
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 奇门仓储字段
	ExpressName string `json:"expressName,omitempty" xml:"expressName,omitempty"`
	// 奇门仓储字段
	BrandCode string `json:"brandCode,omitempty" xml:"brandCode,omitempty"`
	// 奇门仓储字段
	BrandName string `json:"brandName,omitempty" xml:"brandName,omitempty"`
	// 奇门仓储字段
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
