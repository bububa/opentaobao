package normalvisa

import (
	"sync"
)

// QueryVisaApplicantResult 结构体
type QueryVisaApplicantResult struct {
	// 申请人信息列表
	ApplicantInfoList []VisaApplicantInfo `json:"applicant_info_list,omitempty" xml:"applicant_info_list>visa_applicant_info,omitempty"`
	// 总申请人数量
	TotalApplicantsCount int64 `json:"total_applicants_count,omitempty" xml:"total_applicants_count,omitempty"`
}

var poolQueryVisaApplicantResult = sync.Pool{
	New: func() any {
		return new(QueryVisaApplicantResult)
	},
}

// GetQueryVisaApplicantResult() 从对象池中获取QueryVisaApplicantResult
func GetQueryVisaApplicantResult() *QueryVisaApplicantResult {
	return poolQueryVisaApplicantResult.Get().(*QueryVisaApplicantResult)
}

// ReleaseQueryVisaApplicantResult 释放QueryVisaApplicantResult
func ReleaseQueryVisaApplicantResult(v *QueryVisaApplicantResult) {
	v.ApplicantInfoList = v.ApplicantInfoList[:0]
	v.TotalApplicantsCount = 0
	poolQueryVisaApplicantResult.Put(v)
}
