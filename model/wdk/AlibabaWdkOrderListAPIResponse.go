package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderListAPIResponse 五道口订单拉取 API返回值
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
type AlibabaWdkOrderListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderListAPIResponseModel).Reset()
}

// AlibabaWdkOrderListAPIResponseModel is 五道口订单拉取 成功返回结果
type AlibabaWdkOrderListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *AlibabaWdkOrderListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOrderListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderListAPIResponse)
	},
}

// GetAlibabaWdkOrderListAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderListAPIResponse
func GetAlibabaWdkOrderListAPIResponse() *AlibabaWdkOrderListAPIResponse {
	return poolAlibabaWdkOrderListAPIResponse.Get().(*AlibabaWdkOrderListAPIResponse)
}

// ReleaseAlibabaWdkOrderListAPIResponse 将 AlibabaWdkOrderListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderListAPIResponse(v *AlibabaWdkOrderListAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderListAPIResponse.Put(v)
}
