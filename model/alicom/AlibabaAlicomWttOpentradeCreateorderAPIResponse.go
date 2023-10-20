package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomWttOpentradeCreateorderAPIResponse 充值送活动下单接口 API返回值
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
type AlibabaAlicomWttOpentradeCreateorderAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomWttOpentradeCreateorderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomWttOpentradeCreateorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomWttOpentradeCreateorderAPIResponseModel).Reset()
}

// AlibabaAlicomWttOpentradeCreateorderAPIResponseModel is 充值送活动下单接口 成功返回结果
type AlibabaAlicomWttOpentradeCreateorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_wtt_opentrade_createorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomWttOpentradeCreateorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomWttOpentradeCreateorderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomWttOpentradeCreateorderAPIResponse)
	},
}

// GetAlibabaAlicomWttOpentradeCreateorderAPIResponse 从 sync.Pool 获取 AlibabaAlicomWttOpentradeCreateorderAPIResponse
func GetAlibabaAlicomWttOpentradeCreateorderAPIResponse() *AlibabaAlicomWttOpentradeCreateorderAPIResponse {
	return poolAlibabaAlicomWttOpentradeCreateorderAPIResponse.Get().(*AlibabaAlicomWttOpentradeCreateorderAPIResponse)
}

// ReleaseAlibabaAlicomWttOpentradeCreateorderAPIResponse 将 AlibabaAlicomWttOpentradeCreateorderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomWttOpentradeCreateorderAPIResponse(v *AlibabaAlicomWttOpentradeCreateorderAPIResponse) {
	v.Reset()
	poolAlibabaAlicomWttOpentradeCreateorderAPIResponse.Put(v)
}
