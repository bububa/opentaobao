package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeRefundQueryAPIRequest
外部渠道查询退货订单详情接口 API请求
alibaba.wdk.trade.refund.query

该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。 */
type AlibabaWdkTradeRefundQueryAPIRequest struct {
	model.Params
	// 查询请求
	_refundGoodsQueryRequest *RefundGoodsQueryRequest
}

// New
