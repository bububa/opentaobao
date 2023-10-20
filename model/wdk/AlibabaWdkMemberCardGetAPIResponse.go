package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmembercardgetAPIResponse 查询会员卡信息 API返回值
// alibaba.wdk.member.card.get
//
// 根据会员卡查询会员信息
type AlibabawdkmembercardgetAPIResponse struct {
	model.CommonResponse
	AlibabawdkmembercardgetAPIResponseModel
}

// AlibabawdkmembercardgetAPIResponseModel is 查询会员卡信息 成功返回结果
type AlibabawdkmembercardgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_member_card_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabawdkmembercardgetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
