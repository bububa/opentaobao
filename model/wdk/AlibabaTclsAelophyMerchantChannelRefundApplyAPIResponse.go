package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelrefundapplyAPIResponse 翱象商家自有渠道 逆向单申请 API返回值
// alibaba.tcls.aelophy.merchant.channel.refund.apply
//
// 翱象小程序 用户逆向单申请
type AlibabatclsaelophymerchantchannelrefundapplyAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantchannelrefundapplyAPIResponseModel
}

// AlibabatclsaelophymerchantchannelrefundapplyAPIResponseModel is 翱象商家自有渠道 逆向单申请 成功返回结果
type AlibabatclsaelophymerchantchannelrefundapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_refund_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabatclsaelophymerchantchannelrefundapplyApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
