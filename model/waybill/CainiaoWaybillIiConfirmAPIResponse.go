package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiConfirmAPIResponse 物流订单确认接口 API返回值
// cainiao.waybill.ii.confirm
//
// 物流订单确认
type CainiaoWaybillIiConfirmAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiConfirmAPIResponseModel).Reset()
}

// CainiaoWaybillIiConfirmAPIResponseModel is 物流订单确认接口 成功返回结果
type CainiaoWaybillIiConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果
	ConfirmResponse []WaybillOrderConfirmResponse `json:"confirm_response,omitempty" xml:"confirm_response>waybill_order_confirm_response,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ConfirmResponse = m.ConfirmResponse[:0]
}

var poolCainiaoWaybillIiConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiConfirmAPIResponse)
	},
}

// GetCainiaoWaybillIiConfirmAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiConfirmAPIResponse
func GetCainiaoWaybillIiConfirmAPIResponse() *CainiaoWaybillIiConfirmAPIResponse {
	return poolCainiaoWaybillIiConfirmAPIResponse.Get().(*CainiaoWaybillIiConfirmAPIResponse)
}

// ReleaseCainiaoWaybillIiConfirmAPIResponse 将 CainiaoWaybillIiConfirmAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiConfirmAPIResponse(v *CainiaoWaybillIiConfirmAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiConfirmAPIResponse.Put(v)
}
