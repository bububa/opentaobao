package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsConsignTcConfirmAPIRequest
通知交易确认发货接口 API请求
taobao.logistics.consign.tc.confirm

下述业务场景可以使用此接口通知相关的交易订单发货：
1、发货过程分为多段操作，在确定发货之前，不需要通知交易，当货确认已发出之后，才通知交易发货。
2、发货过程涉及到多个订单，其中一个订单是跟实操的发货操作同步的，剩下的订单，需要在实操的订单发货之后，一并通知交易发货。 */
type TaobaoLogisticsConsignTcConfirmAPIRequest struct {
	model.Params
	// 货主id
	_providerId int64
	// ERP的名称
	_appName string
	// 扩展字段 K:V;
	_extendFields string
	// 发货的包裹
	_goodsList []ConfirmConsignGoodsDto
	// 已发货的外部订单号
	_outTradeNo string
	// 卖家id
	_sellerId int64
	// 待发货的淘宝主交易号
	_tradeId int64
}

// New
