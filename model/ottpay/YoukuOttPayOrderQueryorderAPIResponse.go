package ottpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQueryorderAPIResponse
查询订单 API返回值
youku.ott.pay.order.queryorder

通过订单号查询订单信息 */
type YoukuOttPayOrderQueryorderAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderQueryorderAPIResponseModel
}

// YoukuOttPayOrderQueryorderAPIResponseModel is 查询订单 成功返回结果
type YoukuOttPayOrderQueryorderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_queryorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// status
	Data *TvOrderQueryResultDto `json:"data,omitempty" xml:"data,omitempty"`
}
