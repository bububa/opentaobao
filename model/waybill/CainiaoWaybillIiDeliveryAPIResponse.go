package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiDeliveryAPIResponse 派件通知接口 API返回值
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
type CainiaoWaybillIiDeliveryAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiDeliveryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiDeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiDeliveryAPIResponseModel).Reset()
}

// CainiaoWaybillIiDeliveryAPIResponseModel is 派件通知接口 成功返回结果
type CainiaoWaybillIiDeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果
	ConfirmResponse *WaybillOrderConfirmResponse `json:"confirm_response,omitempty" xml:"confirm_response,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiDeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ConfirmResponse = nil
}

var poolCainiaoWaybillIiDeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiDeliveryAPIResponse)
	},
}

// GetCainiaoWaybillIiDeliveryAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiDeliveryAPIResponse
func GetCainiaoWaybillIiDeliveryAPIResponse() *CainiaoWaybillIiDeliveryAPIResponse {
	return poolCainiaoWaybillIiDeliveryAPIResponse.Get().(*CainiaoWaybillIiDeliveryAPIResponse)
}

// ReleaseCainiaoWaybillIiDeliveryAPIResponse 将 CainiaoWaybillIiDeliveryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiDeliveryAPIResponse(v *CainiaoWaybillIiDeliveryAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiDeliveryAPIResponse.Put(v)
}
