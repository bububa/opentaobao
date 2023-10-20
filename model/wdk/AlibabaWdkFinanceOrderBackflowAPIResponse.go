package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFinanceOrderBackflowAPIResponse 财务订单回流 API返回值
// alibaba.wdk.finance.order.backflow
//
// 星巴克拉取财务订单回流数据
type AlibabaWdkFinanceOrderBackflowAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFinanceOrderBackflowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFinanceOrderBackflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFinanceOrderBackflowAPIResponseModel).Reset()
}

// AlibabaWdkFinanceOrderBackflowAPIResponseModel is 财务订单回流 成功返回结果
type AlibabaWdkFinanceOrderBackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_finance_order_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabaWdkFinanceOrderBackflowApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFinanceOrderBackflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkFinanceOrderBackflowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFinanceOrderBackflowAPIResponse)
	},
}

// GetAlibabaWdkFinanceOrderBackflowAPIResponse 从 sync.Pool 获取 AlibabaWdkFinanceOrderBackflowAPIResponse
func GetAlibabaWdkFinanceOrderBackflowAPIResponse() *AlibabaWdkFinanceOrderBackflowAPIResponse {
	return poolAlibabaWdkFinanceOrderBackflowAPIResponse.Get().(*AlibabaWdkFinanceOrderBackflowAPIResponse)
}

// ReleaseAlibabaWdkFinanceOrderBackflowAPIResponse 将 AlibabaWdkFinanceOrderBackflowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFinanceOrderBackflowAPIResponse(v *AlibabaWdkFinanceOrderBackflowAPIResponse) {
	v.Reset()
	poolAlibabaWdkFinanceOrderBackflowAPIResponse.Put(v)
}
