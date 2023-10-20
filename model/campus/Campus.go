package campus

import (
	"sync"
)

// Campus 结构体
type Campus struct {
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// creator
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// modifier
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 地址信息
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 省份编码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 国家编码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 园区编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 园区英文名称
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 园区名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 所属公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 室内面积
	IndoorArea string `json:"indoor_area,omitempty" xml:"indoor_area,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 来源，0自建 1租赁
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 状态，0停用，1启用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 面积
	Area float64 `json:"area,omitempty" xml:"area,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 所属公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 是否删除，0未删除，1删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}

var poolCampus = sync.Pool{
	New: func() any {
		return new(Campus)
	},
}

// GetCampus() 从对象池中获取Campus
func GetCampus() *Campus {
	return poolCampus.Get().(*Campus)
}

// ReleaseCampus 释放Campus
func ReleaseCampus(v *Campus) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Creator = ""
	v.Modifier = ""
	v.Address = ""
	v.CityCode = ""
	v.CityName = ""
	v.ProvinceCode = ""
	v.ProvinceName = ""
	v.CountryCode = ""
	v.CountryName = ""
	v.Code = ""
	v.EnName = ""
	v.Name = ""
	v.CompanyName = ""
	v.IndoorArea = ""
	v.Id = 0
	v.Type = 0
	v.Status = 0
	v.Area = 0
	v.OrderNo = 0
	v.CompanyId = 0
	v.IsDelete = false
	poolCampus.Put(v)
}
