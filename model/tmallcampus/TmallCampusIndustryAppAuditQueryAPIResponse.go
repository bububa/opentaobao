package tmallcampus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCampusIndustryAppAuditQueryAPIResponse 天猫校园查询学生认证信息 API返回值
// tmall.campus.industry.app.audit.query
//
// 天猫校园查询学生认证信息
type TmallCampusIndustryAppAuditQueryAPIResponse struct {
	model.CommonResponse
	TmallCampusIndustryAppAuditQueryAPIResponseModel
}

// TmallCampusIndustryAppAuditQueryAPIResponseModel is 天猫校园查询学生认证信息 成功返回结果
type TmallCampusIndustryAppAuditQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_campus_industry_app_audit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TResult `json:"result,omitempty" xml:"result,omitempty"`
}
