package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest
alibaba.alihealth.baby.baseinfo.order.sync API请求
alibaba.alihealth.baby.baseinfo.order.sync

育学园将订单信息回传给我们 */
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest struct {
	model.Params
	// 健康id
	_tpUserId int64
	// 商品id
	_commodityId string
	// 商品名称
	_commodityName string
	// 价钱
	_amount *BigDecimal
	// 状态，1是已支付，2是已退款
	_status int64
	// 订单时间
	_orderTime string
	// 订单id
	_orderId string
}

// New
