package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaovasordersearchAPIResponse 订单记录导出 API返回值
// taobao.vas.order.search
//
// 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。&lt;br/&gt;建议用于查询前一日的历史记录，不适合用作实时数据查询。&lt;br/&gt;现在只能查询90天以内的数据&lt;br/&gt;该接口限制每分钟所有appkey调用总和只能有800次。
type TaobaovasordersearchAPIResponse struct {
	model.CommonResponse
	TaobaovasordersearchAPIResponseModel
}

// TaobaovasordersearchAPIResponseModel is 订单记录导出 成功返回结果
type TaobaovasordersearchAPIResponseModel struct {
	XMLName xml.Name `xml:"vas_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品订单对象
	ArticleBizOrders []ArticleBizOrder `json:"article_biz_orders,omitempty" xml:"article_biz_orders>article_biz_order,omitempty"`
	// 总记录数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}
