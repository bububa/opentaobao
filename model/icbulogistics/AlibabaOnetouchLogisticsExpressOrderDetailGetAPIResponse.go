package icbulogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse 订单详细信息(面单及仓库信息) API返回值
// alibaba.onetouch.logistics.express.order.detail.get
//
// 订单详细信息(面单及仓库信息)
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponseModel).Reset()
}

// AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponseModel is 订单详细信息(面单及仓库信息) 成功返回结果
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_order_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaOnetouchLogisticsExpressOrderDetailGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse)
	},
}

// GetAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse
func GetAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse() *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse {
	return poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse.Get().(*AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse)
}

// ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse 将 AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse(v *AlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressOrderDetailGetAPIResponse.Put(v)
}
