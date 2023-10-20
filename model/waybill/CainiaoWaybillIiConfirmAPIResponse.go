package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiconfirmAPIResponse 物流订单确认接口 API返回值
// cainiao.waybill.ii.confirm
//
// 物流订单确认
type CainiaowaybilliiconfirmAPIResponse struct {
	model.CommonResponse
	CainiaowaybilliiconfirmAPIResponseModel
}

// CainiaowaybilliiconfirmAPIResponseModel is 物流订单确认接口 成功返回结果
type CainiaowaybilliiconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果
	ConfirmResponse []WaybillOrderConfirmResponse `json:"confirm_response,omitempty" xml:"confirm_response>waybill_order_confirm_response,omitempty"`
}
