package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentOrderGetAPIResponse 闲鱼帮卖订单查询 API返回值
// alibaba.idle.consignment.order.get
//
// 闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
type AlibabaIdleConsignmentOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleConsignmentOrderGetAPIResponseModel).Reset()
}

// AlibabaIdleConsignmentOrderGetAPIResponseModel is 闲鱼帮卖订单查询 成功返回结果
type AlibabaIdleConsignmentOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignment_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleConsignmentOrderGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleConsignmentOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentOrderGetAPIResponse)
	},
}

// GetAlibabaIdleConsignmentOrderGetAPIResponse 从 sync.Pool 获取 AlibabaIdleConsignmentOrderGetAPIResponse
func GetAlibabaIdleConsignmentOrderGetAPIResponse() *AlibabaIdleConsignmentOrderGetAPIResponse {
	return poolAlibabaIdleConsignmentOrderGetAPIResponse.Get().(*AlibabaIdleConsignmentOrderGetAPIResponse)
}

// ReleaseAlibabaIdleConsignmentOrderGetAPIResponse 将 AlibabaIdleConsignmentOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleConsignmentOrderGetAPIResponse(v *AlibabaIdleConsignmentOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaIdleConsignmentOrderGetAPIResponse.Put(v)
}
