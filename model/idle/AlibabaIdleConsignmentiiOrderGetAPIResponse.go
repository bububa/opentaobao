package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentiiOrderGetAPIResponse 闲鱼寄卖V2订单查询 API返回值
// alibaba.idle.consignmentii.order.get
//
// 闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
type AlibabaIdleConsignmentiiOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentiiOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentiiOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleConsignmentiiOrderGetAPIResponseModel).Reset()
}

// AlibabaIdleConsignmentiiOrderGetAPIResponseModel is 闲鱼寄卖V2订单查询 成功返回结果
type AlibabaIdleConsignmentiiOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignmentii_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleConsignmentiiOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentiiOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleConsignmentiiOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentiiOrderGetAPIResponse)
	},
}

// GetAlibabaIdleConsignmentiiOrderGetAPIResponse 从 sync.Pool 获取 AlibabaIdleConsignmentiiOrderGetAPIResponse
func GetAlibabaIdleConsignmentiiOrderGetAPIResponse() *AlibabaIdleConsignmentiiOrderGetAPIResponse {
	return poolAlibabaIdleConsignmentiiOrderGetAPIResponse.Get().(*AlibabaIdleConsignmentiiOrderGetAPIResponse)
}

// ReleaseAlibabaIdleConsignmentiiOrderGetAPIResponse 将 AlibabaIdleConsignmentiiOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleConsignmentiiOrderGetAPIResponse(v *AlibabaIdleConsignmentiiOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaIdleConsignmentiiOrderGetAPIResponse.Put(v)
}
