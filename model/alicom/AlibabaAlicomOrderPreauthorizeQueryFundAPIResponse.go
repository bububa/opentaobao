package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse 资金流水查询 API返回值
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
type AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomOrderPreauthorizeQueryFundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomOrderPreauthorizeQueryFundAPIResponseModel).Reset()
}

// AlibabaAlicomOrderPreauthorizeQueryFundAPIResponseModel is 资金流水查询 成功返回结果
type AlibabaAlicomOrderPreauthorizeQueryFundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_query_fund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderPreauthorizeQueryFundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse)
	},
}

// GetAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse 从 sync.Pool 获取 AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse
func GetAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse() *AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse {
	return poolAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse.Get().(*AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse)
}

// ReleaseAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse 将 AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse(v *AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse) {
	v.Reset()
	poolAlibabaAlicomOrderPreauthorizeQueryFundAPIResponse.Put(v)
}
