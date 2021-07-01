package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeConsigngoodsAPIRequest
卖家发货 API请求
tmall.exchange.consigngoods

卖家发货 */
type TmallExchangeConsigngoodsAPIRequest struct {
	model.Params
	// 换货单号ID
	_disputeId int64
	// 卖家发货的物流单号
	_logisticsNo string
	// 卖家发货的物流类型，100表示平邮，200表示快递
	_logisticsType int64
	// 卖家发货的快递公司
	_logisticsCompanyName string
	// 返回字段
	_fields []string
}

// New
