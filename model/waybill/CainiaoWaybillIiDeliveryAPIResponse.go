package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliideliveryAPIResponse 派件通知接口 API返回值
// cainiao.waybill.ii.delivery
//
// 极效前置场景下的使用此接口，通知进行派件
type CainiaowaybilliideliveryAPIResponse struct {
	model.CommonResponse
	CainiaowaybilliideliveryAPIResponseModel
}

// CainiaowaybilliideliveryAPIResponseModel is 派件通知接口 成功返回结果
type CainiaowaybilliideliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 确认结果
	ConfirmResponse *WaybillOrderConfirmResponse `json:"confirm_response,omitempty" xml:"confirm_response,omitempty"`
}
