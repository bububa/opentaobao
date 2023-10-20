package ascp

import (
	"sync"
)

// SiteUpsetRequest 结构体
type SiteUpsetRequest struct {
	// 网点小件员信息
	CourierInfos []CourierInfos `json:"courier_infos,omitempty" xml:"courier_infos>courier_infos,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 网点编码
	SiteCode string `json:"site_code,omitempty" xml:"site_code,omitempty"`
	// 网点名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 网点所在的行政地址id（菜鸟地址库id，三级/四级地址id
	AddressId string `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 网点联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 网点联系人手机号
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 网点固定电话
	ContactTel string `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 网点状态  1-启用、2-停用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSiteUpsetRequest = sync.Pool{
	New: func() any {
		return new(SiteUpsetRequest)
	},
}

// GetSiteUpsetRequest() 从对象池中获取SiteUpsetRequest
func GetSiteUpsetRequest() *SiteUpsetRequest {
	return poolSiteUpsetRequest.Get().(*SiteUpsetRequest)
}

// ReleaseSiteUpsetRequest 释放SiteUpsetRequest
func ReleaseSiteUpsetRequest(v *SiteUpsetRequest) {
	v.CourierInfos = v.CourierInfos[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.SiteCode = ""
	v.SiteName = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.Address = ""
	v.AddressId = ""
	v.ContactName = ""
	v.ContactMobile = ""
	v.ContactTel = ""
	v.RequestTime = 0
	v.Status = 0
	poolSiteUpsetRequest.Put(v)
}
