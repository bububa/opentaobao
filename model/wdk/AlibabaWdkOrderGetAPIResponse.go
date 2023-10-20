package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderGetAPIResponse 交易订单详情查询 API返回值
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
type AlibabaWdkOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderGetAPIResponseModel).Reset()
}

// AlibabaWdkOrderGetAPIResponseModel is 交易订单详情查询 成功返回结果
type AlibabaWdkOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *AlibabaWdkOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderGetAPIResponse)
	},
}

// GetAlibabaWdkOrderGetAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderGetAPIResponse
func GetAlibabaWdkOrderGetAPIResponse() *AlibabaWdkOrderGetAPIResponse {
	return poolAlibabaWdkOrderGetAPIResponse.Get().(*AlibabaWdkOrderGetAPIResponse)
}

// ReleaseAlibabaWdkOrderGetAPIResponse 将 AlibabaWdkOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderGetAPIResponse(v *AlibabaWdkOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderGetAPIResponse.Put(v)
}
