package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderconfirmAPIResponse 发货单确认接口 API返回值
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
type TaobaoqimendeliveryorderconfirmAPIResponse struct {
	model.CommonResponse
	TaobaoqimendeliveryorderconfirmAPIResponseModel
}

// TaobaoqimendeliveryorderconfirmAPIResponseModel is 发货单确认接口 成功返回结果
type TaobaoqimendeliveryorderconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimendeliveryorderconfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}
