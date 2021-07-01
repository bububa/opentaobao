package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderConfirmAPIResponse
发货单确认接口 API返回值
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm */
type TaobaoQimenDeliveryorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderConfirmAPIResponseModel
}

// TaobaoQimenDeliveryorderConfirmAPIResponseModel is 发货单确认接口 成功返回结果
type TaobaoQimenDeliveryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}
