package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchconfirmAPIResponse 发货单确认接口 API返回值
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel
}

// TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel is 发货单确认接口 成功返回结果
type TaobaoQimenDeliveryorderBatchconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_batchconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenDeliveryorderBatchconfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}
