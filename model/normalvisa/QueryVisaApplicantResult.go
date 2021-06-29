package normalvisa

// QueryVisaApplicantResult 
type QueryVisaApplicantResult struct {
    // 申请人信息列表
    ApplicantInfoList   []VisaApplicantInfo `json:"applicant_info_list,omitempty" xml:"applicant_info_list>visa_applicant_info,omitempty"`
    // 总申请人数量
    TotalApplicantsCount   int64 `json:"total_applicants_count,omitempty" xml:"total_applicants_count,omitempty"`
}
