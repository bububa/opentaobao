package normalvisa

import (
	"sync"
)

// QueryApplicantParam 结构体
type QueryApplicantParam struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 页面大小，默认20，最大支持的页面大小为500。
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页，默认第一页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 查询类型，默认为1。1-按用户提交材料时间查询（仅查询状态为1003的申请人）
	QueryType int64 `json:"query_type,omitempty" xml:"query_type,omitempty"`
}

var poolQueryApplicantParam = sync.Pool{
	New: func() any {
		return new(QueryApplicantParam)
	},
}

// GetQueryApplicantParam() 从对象池中获取QueryApplicantParam
func GetQueryApplicantParam() *QueryApplicantParam {
	return poolQueryApplicantParam.Get().(*QueryApplicantParam)
}

// ReleaseQueryApplicantParam 释放QueryApplicantParam
func ReleaseQueryApplicantParam(v *QueryApplicantParam) {
	v.StartTime = ""
	v.EndTime = ""
	v.PageSize = 0
	v.CurrentPage = 0
	v.QueryType = 0
	poolQueryApplicantParam.Put(v)
}
