package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkopenOrderGetAPIResponse 五道口商户订单获取 API返回值
// alibaba.wdkopen.order.get
//
// 商户通过五道口订单id获取订单信息
type AlibabaWdkopenOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkopenOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkopenOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkopenOrderGetAPIResponseModel).Reset()
}

// AlibabaWdkopenOrderGetAPIResponseModel is 五道口商户订单获取 成功返回结果
type AlibabaWdkopenOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkopen_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	TopBaseResult *TopBaseResult `json:"top_base_result,omitempty" xml:"top_base_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkopenOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopBaseResult = nil
}

var poolAlibabaWdkopenOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkopenOrderGetAPIResponse)
	},
}

// GetAlibabaWdkopenOrderGetAPIResponse 从 sync.Pool 获取 AlibabaWdkopenOrderGetAPIResponse
func GetAlibabaWdkopenOrderGetAPIResponse() *AlibabaWdkopenOrderGetAPIResponse {
	return poolAlibabaWdkopenOrderGetAPIResponse.Get().(*AlibabaWdkopenOrderGetAPIResponse)
}

// ReleaseAlibabaWdkopenOrderGetAPIResponse 将 AlibabaWdkopenOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkopenOrderGetAPIResponse(v *AlibabaWdkopenOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkopenOrderGetAPIResponse.Put(v)
}
