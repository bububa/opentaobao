package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse 获取运力时间片信息 API返回值
// alibaba.tcls.aelophy.merchant.channel.order.sliceget
//
// 获取履约时间片
type AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel
}

// AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel is 获取运力时间片信息 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_sliceget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderSlicegetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
