package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsReturnBillGetAPIRequest
获取销退收货信息 API请求
taobao.wlb.wms.return.bill.get

通过订单号获取单个销退单收货信息。 */
type TaobaoWlbWmsReturnBillGetAPIRequest struct {
	model.Params
	// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_orderCode string
	// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_cnOrderCode string
}

// New
