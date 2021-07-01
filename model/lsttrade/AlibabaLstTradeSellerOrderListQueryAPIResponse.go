package lsttrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOrderListQueryAPIResponse
订单列表查看(卖家视角) API返回值
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表 */
type AlibabaLstTradeSellerOrderListQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeSellerOrderListQueryAPIResponseModel
}

// AlibabaLstTradeSellerOrderListQueryAPIResponseModel is 订单列表查看(卖家视角) 成功返回结果
type AlibabaLstTradeSellerOrderListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstTradeSellerOrderListQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
