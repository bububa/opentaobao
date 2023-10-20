package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderQueryAPIResponse 发货单查询接口 API返回值
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenDeliveryorderQueryAPIResponseModel
}

// TaobaoQimenDeliveryorderQueryAPIResponseModel is 发货单查询接口 成功返回结果
type TaobaoQimenDeliveryorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_deliveryorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *DeliveryOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
