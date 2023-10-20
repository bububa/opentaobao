package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderprocessqueryAPIResponse 订单流水查询接口 API返回值
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
type TaobaoqimenorderprocessqueryAPIResponse struct {
	model.CommonResponse
	TaobaoqimenorderprocessqueryAPIResponseModel
}

// TaobaoqimenorderprocessqueryAPIResponseModel is 订单流水查询接口 成功返回结果
type TaobaoqimenorderprocessqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderprocess_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *OrderProcessQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
