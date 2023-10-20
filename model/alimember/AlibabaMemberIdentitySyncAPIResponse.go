package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberidentitysyncAPIResponse 会员身份信息同步 API返回值
// alibaba.member.identity.sync
//
// 会员身份信息同步
type AlibabamemberidentitysyncAPIResponse struct {
	model.CommonResponse
	AlibabamemberidentitysyncAPIResponseModel
}

// AlibabamemberidentitysyncAPIResponseModel is 会员身份信息同步 成功返回结果
type AlibabamemberidentitysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_identity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabamemberidentitysyncTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
