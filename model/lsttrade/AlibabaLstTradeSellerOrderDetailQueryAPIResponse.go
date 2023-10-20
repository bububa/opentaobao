package lsttrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerorderdetailqueryAPIResponse 订单详情查看(卖家视角) API返回值
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
type AlibabalsttradesellerorderdetailqueryAPIResponse struct {
	model.CommonResponse
	AlibabalsttradesellerorderdetailqueryAPIResponseModel
}

// AlibabalsttradesellerorderdetailqueryAPIResponseModel is 订单详情查看(卖家视角) 成功返回结果
type AlibabalsttradesellerorderdetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabalsttradesellerorderdetailqueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
