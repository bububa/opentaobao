package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberIdentityRescindfinishAPIResponse
取消确认 API返回值
alibaba.member.identity.rescindfinish

取消确认 */
type AlibabaMemberIdentityRescindfinishAPIResponse struct {
	model.CommonResponse
	AlibabaMemberIdentityRescindfinishAPIResponseModel
}

// AlibabaMemberIdentityRescindfinishAPIResponseModel is 取消确认 成功返回结果
type AlibabaMemberIdentityRescindfinishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_identity_rescindfinish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}
