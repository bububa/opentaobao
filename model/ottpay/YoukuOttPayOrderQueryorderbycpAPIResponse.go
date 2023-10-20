package ottpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryorderbycpAPIResponse 订单查询接口(cp订单号查询) API返回值
// youku.ott.pay.order.queryorderbycp
//
// 给商户服务端查询订单状态
type YoukuOttPayOrderQueryorderbycpAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderQueryorderbycpAPIResponseModel
}

// YoukuOttPayOrderQueryorderbycpAPIResponseModel is 订单查询接口(cp订单号查询) 成功返回结果
type YoukuOttPayOrderQueryorderbycpAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_queryorderbycp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderQueryResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
