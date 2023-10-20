package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse 计算快递运费&下单参数校验 API返回值
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponseModel is 计算快递运费&下单参数校验 成功返回结果
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_charge_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressChargeCalculateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse
func GetAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse() *AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse 将 AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse(v *AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse.Put(v)
}
