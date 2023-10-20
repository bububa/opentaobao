package flight

import (
	"sync"
)

// PolicyQueryParamDto 结构体
type PolicyQueryParamDto struct {
	// 航司二字码
	Airline string `json:"airline,omitempty" xml:"airline,omitempty"`
	// 到达机场三字码
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 去程舱位
	CabinList1 string `json:"cabin_list1,omitempty" xml:"cabin_list1,omitempty"`
	// 回程舱位
	CabinList2 string `json:"cabin_list2,omitempty" xml:"cabin_list2,omitempty"`
	// 出发机场三字码
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 政策编码
	PolicyCode string `json:"policy_code,omitempty" xml:"policy_code,omitempty"`
	// 销售起始日期，不需要传时分秒
	SaleStartDate string `json:"sale_start_date,omitempty" xml:"sale_start_date,omitempty"`
	// 销售截止日期，不需要传时分秒
	SalesEndDate string `json:"sales_end_date,omitempty" xml:"sales_end_date,omitempty"`
	// 行程截止日期，不需要传时分秒
	TravelEndDate string `json:"travel_end_date,omitempty" xml:"travel_end_date,omitempty"`
	// 行程起始日期，不需要传时分秒
	TravelStartDate string `json:"travel_start_date,omitempty" xml:"travel_start_date,omitempty"`
	// 大客户编码，请输入包含PAT在内的完整格式。
	AccountCode string `json:"account_code,omitempty" xml:"account_code,omitempty"`
	// 政策来源：0，手工政策；1，excel政策；2，api政策
	PolicySource int64 `json:"policy_source,omitempty" xml:"policy_source,omitempty"`
	// 政策状态：1，有效；2，挂起；0，删除 3，停用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 根据政策编码模糊查询，模糊查询执行like，暂不支持前缀索引
	FuzzyQuery bool `json:"fuzzy_query,omitempty" xml:"fuzzy_query,omitempty"`
	// 中转运价标识
	IsOwEoe bool `json:"is_ow_eoe,omitempty" xml:"is_ow_eoe,omitempty"`
}

var poolPolicyQueryParamDto = sync.Pool{
	New: func() any {
		return new(PolicyQueryParamDto)
	},
}

// GetPolicyQueryParamDto() 从对象池中获取PolicyQueryParamDto
func GetPolicyQueryParamDto() *PolicyQueryParamDto {
	return poolPolicyQueryParamDto.Get().(*PolicyQueryParamDto)
}

// ReleasePolicyQueryParamDto 释放PolicyQueryParamDto
func ReleasePolicyQueryParamDto(v *PolicyQueryParamDto) {
	v.Airline = ""
	v.ArrAirport = ""
	v.CabinList1 = ""
	v.CabinList2 = ""
	v.DepAirport = ""
	v.PolicyCode = ""
	v.SaleStartDate = ""
	v.SalesEndDate = ""
	v.TravelEndDate = ""
	v.TravelStartDate = ""
	v.AccountCode = ""
	v.PolicySource = 0
	v.Status = 0
	v.FuzzyQuery = false
	v.IsOwEoe = false
	poolPolicyQueryParamDto.Put(v)
}
