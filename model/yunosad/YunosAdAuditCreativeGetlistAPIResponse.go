package yunosad

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosAdAuditCreativeGetlistAPIResponse 批量获取创意审核状态 API返回值
// yunos.ad.audit.creative.getlist
//
// 批量获取创意审核状态
type YunosAdAuditCreativeGetlistAPIResponse struct {
	model.CommonResponse
	YunosAdAuditCreativeGetlistAPIResponseModel
}

// YunosAdAuditCreativeGetlistAPIResponseModel is 批量获取创意审核状态 成功返回结果
type YunosAdAuditCreativeGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_ad_audit_creative_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []CreativeAuditDto `json:"results,omitempty" xml:"results>creative_audit_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// isOk
	IsOk string `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
	// errorCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
