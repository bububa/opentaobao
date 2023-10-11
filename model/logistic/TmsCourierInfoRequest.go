package logistic

// TmsCourierInfoRequest 结构体
type TmsCourierInfoRequest struct {
	// 小件员名称
	CourierName string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
	// 小件员电话号码
	CourierMobile string `json:"courier_mobile,omitempty" xml:"courier_mobile,omitempty"`
	// 小件员编码
	CourierNo string `json:"courier_no,omitempty" xml:"courier_no,omitempty"`
	// 小件员所属的网点名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 小件员所属的网点编码
	SiteCode string `json:"site_code,omitempty" xml:"site_code,omitempty"`
}
