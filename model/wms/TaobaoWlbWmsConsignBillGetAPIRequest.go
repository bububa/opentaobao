package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsConsignBillGetAPIRequest
获取销售订单发货信息 API请求
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息 */
type TaobaoWlbWmsConsignBillGetAPIRequest struct {
	model.Params
	// 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
	_cnOrderCode string
	// ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
	_orderCode string
}

// New
