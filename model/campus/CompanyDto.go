package campus

// CompanyDTO 
type CompanyDTO struct {
    // 园区名称
    CampusName   string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
    // 园区ID
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    // 公司联系电话
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 公司状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 是否默认公司
    IsDefault   bool `json:"is_default,omitempty" xml:"is_default,omitempty"`
    // 公司简称
    ShortName   string `json:"short_name,omitempty" xml:"short_name,omitempty"`
    // 公司名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 公司ID
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    // 公司ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 是否物业公司
    IsWuye   bool `json:"is_wuye,omitempty" xml:"is_wuye,omitempty"`
    // 签约主体公司id
    HrSignCompanyId   string `json:"hr_sign_company_id,omitempty" xml:"hr_sign_company_id,omitempty"`
    // corpId
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    // 公司编号
    CompanyCode   string `json:"company_code,omitempty" xml:"company_code,omitempty"`
    // 公司人数
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
}
