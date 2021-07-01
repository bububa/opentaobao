package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderstatusUpdateAPIRequest
订单状态更新接口 API请求
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。 */
type TaobaoQimenOrderstatusUpdateAPIRequest struct {
	model.Params
	// 星盘派单号
	_allocationCode string
	// 淘系子订单号
	_orderCodes []int64
	// 目标门店的商户中心门店编码
	_storeId int64
	// 业务类型，（枚举值：FAHUO、ZITI）
	_type string
	// 订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY
	_status string
	// 操作人
	_operator string
	// 事件发生时间
	_actionTime string
}

// New
