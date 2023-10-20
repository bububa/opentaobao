package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessSupplyChargeAPIResponse 供销商品充值接口 API返回值
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
type AlibabaAiContentBusinessSupplyChargeAPIResponse struct {
	model.CommonResponse
	AlibabaAiContentBusinessSupplyChargeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAiContentBusinessSupplyChargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAiContentBusinessSupplyChargeAPIResponseModel).Reset()
}

// AlibabaAiContentBusinessSupplyChargeAPIResponseModel is 供销商品充值接口 成功返回结果
type AlibabaAiContentBusinessSupplyChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_content_business_supply_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAiContentBusinessSupplyChargeBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAiContentBusinessSupplyChargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAiContentBusinessSupplyChargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAiContentBusinessSupplyChargeAPIResponse)
	},
}

// GetAlibabaAiContentBusinessSupplyChargeAPIResponse 从 sync.Pool 获取 AlibabaAiContentBusinessSupplyChargeAPIResponse
func GetAlibabaAiContentBusinessSupplyChargeAPIResponse() *AlibabaAiContentBusinessSupplyChargeAPIResponse {
	return poolAlibabaAiContentBusinessSupplyChargeAPIResponse.Get().(*AlibabaAiContentBusinessSupplyChargeAPIResponse)
}

// ReleaseAlibabaAiContentBusinessSupplyChargeAPIResponse 将 AlibabaAiContentBusinessSupplyChargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAiContentBusinessSupplyChargeAPIResponse(v *AlibabaAiContentBusinessSupplyChargeAPIResponse) {
	v.Reset()
	poolAlibabaAiContentBusinessSupplyChargeAPIResponse.Put(v)
}
