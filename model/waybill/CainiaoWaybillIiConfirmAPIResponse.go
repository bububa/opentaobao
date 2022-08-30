package waybill

import (
	"encoding/xml"

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

// CainiaoWaybillIiConfirmAPIResponseModel is 物流订单确认接口 成功返回结果
type CainiaoWaybillIiConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果
	ConfirmResponse []WaybillOrderConfirmResponse `json:"confirm_response,omitempty" xml:"confirm_response>waybill_order_confirm_response,omitempty"`
}
