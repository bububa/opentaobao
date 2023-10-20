package consignplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoConsignplatformOrderCancelAPIResponse 菜鸟发货工作台取消包裹以及订单 API返回值
// cainiao.consignplatform.order.cancel
//
// 菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。
type CainiaoConsignplatformOrderCancelAPIResponse struct {
	model.CommonResponse
	CainiaoConsignplatformOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoConsignplatformOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoConsignplatformOrderCancelAPIResponseModel).Reset()
}

// CainiaoConsignplatformOrderCancelAPIResponseModel is 菜鸟发货工作台取消包裹以及订单 成功返回结果
type CainiaoConsignplatformOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_consignplatform_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败信息
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 失败code
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// 取消是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoConsignplatformOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailCode = ""
	m.FailMessage = ""
	m.Result = false
}

var poolCainiaoConsignplatformOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoConsignplatformOrderCancelAPIResponse)
	},
}

// GetCainiaoConsignplatformOrderCancelAPIResponse 从 sync.Pool 获取 CainiaoConsignplatformOrderCancelAPIResponse
func GetCainiaoConsignplatformOrderCancelAPIResponse() *CainiaoConsignplatformOrderCancelAPIResponse {
	return poolCainiaoConsignplatformOrderCancelAPIResponse.Get().(*CainiaoConsignplatformOrderCancelAPIResponse)
}

// ReleaseCainiaoConsignplatformOrderCancelAPIResponse 将 CainiaoConsignplatformOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseCainiaoConsignplatformOrderCancelAPIResponse(v *CainiaoConsignplatformOrderCancelAPIResponse) {
	v.Reset()
	poolCainiaoConsignplatformOrderCancelAPIResponse.Put(v)
}
