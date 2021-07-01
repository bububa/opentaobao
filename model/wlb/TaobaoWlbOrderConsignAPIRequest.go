package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderConsignAPIRequest
物流宝订单已发货通知接口 API请求
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
type TaobaoWlbOrderConsignAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// New
