package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse 快递下单 API返回值
// alibaba.onetouch.logistics.express.logistics.order.create
//
// 快递下单
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponseModel is 快递下单 成功返回结果
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_logistics_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse
func GetAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse 将 AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse(v *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIResponse.Put(v)
}
