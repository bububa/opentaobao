package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenordercancelAPIResponse 单据取消接口 API返回值
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
type TaobaoqimenordercancelAPIResponse struct {
	model.CommonResponse
	TaobaoqimenordercancelAPIResponseModel
}

// TaobaoqimenordercancelAPIResponseModel is 单据取消接口 成功返回结果
type TaobaoqimenordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *OrderCancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}
