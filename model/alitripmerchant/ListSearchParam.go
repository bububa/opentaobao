package alitripmerchant

import (
	"sync"
)

// ListSearchParam 结构体
type ListSearchParam struct {
	// 星级筛选
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 城市编码，默认上海
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 儿童年龄
	ChildrenAges string `json:"children_ages,omitempty" xml:"children_ages,omitempty"`
	// 入店时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 品牌编码
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 用户登录信息
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 会员等级
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 套餐代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 最高价格
	PriceMax int64 `json:"price_max,omitempty" xml:"price_max,omitempty"`
	// 成人数
	AdultNum int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// 儿童数
	ChildNum int64 `json:"child_num,omitempty" xml:"child_num,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 升降序 0降序 1升序
	Dir int64 `json:"dir,omitempty" xml:"dir,omitempty"`
	// 最低价格
	PriceMin int64 `json:"price_min,omitempty" xml:"price_min,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolListSearchParam = sync.Pool{
	New: func() any {
		return new(ListSearchParam)
	},
}

// GetListSearchParam() 从对象池中获取ListSearchParam
func GetListSearchParam() *ListSearchParam {
	return poolListSearchParam.Get().(*ListSearchParam)
}

// ReleaseListSearchParam 释放ListSearchParam
func ReleaseListSearchParam(v *ListSearchParam) {
	v.Star = ""
	v.CityCode = ""
	v.ChildrenAges = ""
	v.CheckIn = ""
	v.CheckOut = ""
	v.Brand = ""
	v.Token = ""
	v.MemberLevel = ""
	v.VoucherId = ""
	v.PriceMax = 0
	v.AdultNum = 0
	v.ChildNum = 0
	v.Offset = 0
	v.PageSize = 0
	v.Dir = 0
	v.PriceMin = 0
	v.PageNo = 0
	v.Version = 0
	poolListSearchParam.Put(v)
}
