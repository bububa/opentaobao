package legalcase

import (
	"sync"
)

// FeedbackRequestModel 结构体
type FeedbackRequestModel struct {
	// 答辩口径描述
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 类型,法宝主动查询为fabao_query
	FeedBackType string `json:"feed_back_type,omitempty" xml:"feed_back_type,omitempty"`
	// 采纳理由
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 提交人
	SubmitPeople string `json:"submit_people,omitempty" xml:"submit_people,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 口径id
	StandpointId int64 `json:"standpoint_id,omitempty" xml:"standpoint_id,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 观点版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否采纳,主动查询反馈必须为true
	AcceptFlag bool `json:"accept_flag,omitempty" xml:"accept_flag,omitempty"`
}

var poolFeedbackRequestModel = sync.Pool{
	New: func() any {
		return new(FeedbackRequestModel)
	},
}

// GetFeedbackRequestModel() 从对象池中获取FeedbackRequestModel
func GetFeedbackRequestModel() *FeedbackRequestModel {
	return poolFeedbackRequestModel.Get().(*FeedbackRequestModel)
}

// ReleaseFeedbackRequestModel 释放FeedbackRequestModel
func ReleaseFeedbackRequestModel(v *FeedbackRequestModel) {
	v.DefenseCaliber = ""
	v.FeedBackType = ""
	v.Reason = ""
	v.SubmitPeople = ""
	v.EntrustId = 0
	v.StandpointId = 0
	v.SuitId = 0
	v.Version = 0
	v.AcceptFlag = false
	poolFeedbackRequestModel.Put(v)
}
