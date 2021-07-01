package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsStockInBillGetAPIRequest
获取入库单信息 API请求
taobao.wlb.wms.stock.in.bill.get

获取入库单信息 */
type TaobaoWlbWmsStockInBillGetAPIRequest struct {
	model.Params
	// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_orderCode string
	// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_cnOrderCode string
}

// New
