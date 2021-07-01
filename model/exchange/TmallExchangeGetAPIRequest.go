package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeGetAPIRequest
获取单笔换货详情 API请求
tmall.exchange.get

获取单笔换货详情 */
type TmallExchangeGetAPIRequest struct {
	model.Params
	// 换货单号ID
	_disputeId int64
	// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
	_fields []string
}

// New
