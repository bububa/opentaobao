package promotion

import (
	"sync"
)

// CouponTemplateItemQueryRequest 结构体
type CouponTemplateItemQueryRequest struct {
	// 模板表主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 分页信息
	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
	// 分组序号
	LogicGroupNumber int64 `json:"logic_group_number,omitempty" xml:"logic_group_number,omitempty"`
	// 五道口分组id
	WdkGroupId int64 `json:"wdk_group_id,omitempty" xml:"wdk_group_id,omitempty"`
}

var poolCouponTemplateItemQueryRequest = sync.Pool{
	New: func() any {
		return new(CouponTemplateItemQueryRequest)
	},
}

// GetCouponTemplateItemQueryRequest() 从对象池中获取CouponTemplateItemQueryRequest
func GetCouponTemplateItemQueryRequest() *CouponTemplateItemQueryRequest {
	return poolCouponTemplateItemQueryRequest.Get().(*CouponTemplateItemQueryRequest)
}

// ReleaseCouponTemplateItemQueryRequest 释放CouponTemplateItemQueryRequest
func ReleaseCouponTemplateItemQueryRequest(v *CouponTemplateItemQueryRequest) {
	v.Id = 0
	v.PageInfo = nil
	v.SourceId = 0
	v.UserInfo = nil
	v.LogicGroupNumber = 0
	v.WdkGroupId = 0
	poolCouponTemplateItemQueryRequest.Put(v)
}
