package wdk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel).Reset()
}

// AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel is 获取运力时间片信息 成功返回结果
type AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_sliceget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *AlibabaTclsAelophyMerchantChannelOrderSlicegetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse)
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse
func GetAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse() *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse {
	return poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse.Get().(*AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse 将 AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse(v *AlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderSlicegetAPIResponse.Put(v)
}
