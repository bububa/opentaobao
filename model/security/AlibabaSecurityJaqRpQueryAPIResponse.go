package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpqueryAPIResponse 聚安全实人认证查询认证结果 API返回值
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
type AlibabasecurityjaqrpqueryAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqrpqueryAPIResponseModel
}

// AlibabasecurityjaqrpqueryAPIResponseModel is 聚安全实人认证查询认证结果 成功返回结果
type AlibabasecurityjaqrpqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_rp_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果信息
	Data *RpAuditResultBo `json:"data,omitempty" xml:"data,omitempty"`
}
