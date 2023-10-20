package btrip

import (
	"sync"
)

// OpenIsvApplyRs 结构体
type OpenIsvApplyRs struct {
	// 审批人列表
	ApproverList []OpenApproverInfo `json:"approver_list,omitempty" xml:"approver_list>open_approver_info,omitempty"`
	// 行程列表
	ItineraryList []OpenItineraryInfo `json:"itinerary_list,omitempty" xml:"itinerary_list>open_itinerary_info,omitempty"`
	// 出行人列表
	TravelerList []OpenUserInfo `json:"traveler_list,omitempty" xml:"traveler_list>open_user_info,omitempty"`
	// 外部出行人列表
	ExternalTravelerList []OpenUserInfo `json:"external_traveler_list,omitempty" xml:"external_traveler_list>open_user_info,omitempty"`
	// 城市集行程列表
	ItinerarySetList []OpenItineraryInfo `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list>open_itinerary_info,omitempty"`
	// 商旅审批展示id
	ApplyShowId string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	// 商旅企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 企业名称
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 申请人部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 审批状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 第三方业务id
	ThirdpartBusinessId string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// 第三方审批单id,如果非第三方审批单则为空
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 出差事由
	TripCause string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// 审批单标题
	TripTitle string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// 申请人id（第三方用户Id）
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 第三方关联单号
	UnionNo string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// 流程编码
	FlowCode string `json:"flow_code,omitempty" xml:"flow_code,omitempty"`
	// 第三方企业id
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 审批单扩展信息
	ExtendValue string `json:"extend_value,omitempty" xml:"extend_value,omitempty"`
	// 申请理由
	Cause string `json:"cause,omitempty" xml:"cause,omitempty"`
	// 商旅审批单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 审批状态：0审批中 1已同意 2已拒绝 3已转交，4已取消 5已终止 6发起审批 7评论
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差天数
	TripDay int64 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// 申请单提交类型 1代提交 2本人提交 注意：当申请单为代提交时，申请单提交人自己无法为自己下单
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 审批单酒店预算，单位分
	HotelBudget int64 `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	// 审批单机票预算，单位分
	FlightBudget int64 `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	// 审批单火车票预算，单位分
	TrainBudget int64 `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	// 审批单用车预算，单位分
	VehicleBudget int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
	// 审批单总预算，单位分
	Budget int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// 多个申请单预算合并。1：否，【union_no】相同的【申请单(apply id)】，每个的【预算】仅对本申请单生效。2：是，所有【union_no】相同的【申请单(apply id)】，其中全部【预算】合并求和，可以混用。
	BudgetMerge int64 `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	// 0：不限制出行人，1：限申请单内的出行人。注意：不限出行人，实际出行人也不限制差标
	LimitTraveler int64 `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	// 同时预订(机票&amp;火车票)规则。1：就高；2：就低。
	TogetherBookRule int64 `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	// 酒店合住规则
	HotelShare *HotelShareInfo `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	// 申请单城市规则： 0出发&amp;目的地一对一，按列表传行程  1多选N个地点，城市集行程 不传默认为0 会根据商旅管理后台-通用差旅设置-行程城市规则中的设置，校验申请单本字段的值是否正确 当行程城市规则中设置的是“1对1行程”时，必须传0 当行程城市规则中设置的是“多对多城市集行程”时，必须传1 会根据此字段传入的值，校验行程传参是否正确 当申请单城市规则为0，itinerary_list行程列表必填 当申请单城市规则为1，城市集行程必填
	ItineraryRule int64 `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	// 商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 审批类型（5:机票改签审批、6:机票退票审批、3:超标审批）
	BizCategory int64 `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
}

var poolOpenIsvApplyRs = sync.Pool{
	New: func() any {
		return new(OpenIsvApplyRs)
	},
}

// GetOpenIsvApplyRs() 从对象池中获取OpenIsvApplyRs
func GetOpenIsvApplyRs() *OpenIsvApplyRs {
	return poolOpenIsvApplyRs.Get().(*OpenIsvApplyRs)
}

// ReleaseOpenIsvApplyRs 释放OpenIsvApplyRs
func ReleaseOpenIsvApplyRs(v *OpenIsvApplyRs) {
	v.ApproverList = v.ApproverList[:0]
	v.ItineraryList = v.ItineraryList[:0]
	v.TravelerList = v.TravelerList[:0]
	v.ExternalTravelerList = v.ExternalTravelerList[:0]
	v.ItinerarySetList = v.ItinerarySetList[:0]
	v.ApplyShowId = ""
	v.CorpId = ""
	v.CorpName = ""
	v.DepartId = ""
	v.DepartName = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.StatusDesc = ""
	v.ThirdpartBusinessId = ""
	v.ThirdpartId = ""
	v.TripCause = ""
	v.TripTitle = ""
	v.UserId = ""
	v.UserName = ""
	v.UnionNo = ""
	v.FlowCode = ""
	v.ThirdpartCorpId = ""
	v.ExtendValue = ""
	v.Cause = ""
	v.Id = 0
	v.Status = 0
	v.TripDay = 0
	v.Type = 0
	v.HotelBudget = 0
	v.FlightBudget = 0
	v.TrainBudget = 0
	v.VehicleBudget = 0
	v.Budget = 0
	v.BudgetMerge = 0
	v.LimitTraveler = 0
	v.TogetherBookRule = 0
	v.HotelShare = nil
	v.ItineraryRule = 0
	v.ApplyId = 0
	v.BizCategory = 0
	poolOpenIsvApplyRs.Put(v)
}
