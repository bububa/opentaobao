package yunosad

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdAuditCreativeGetAPIResponse
获取单个创意审核状态 API返回值
yunos.ad.audit.creative.get

获取单个创意审核状态 */
type YunosAdAuditCreativeGetAPIResponse struct {
	model.CommonResponse
	YunosAdAuditCreativeGetAPIResponseModel
}

// YunosAdAuditCreativeGetAPIResponseModel is 获取单个创意审核状态 成功返回结果
type YunosAdAuditCreativeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_ad_audit_creative_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 是否成功
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
	// 审核结果
	Result *CreativeAuditDto `json:"result,omitempty" xml:"result,omitempty"`
}
