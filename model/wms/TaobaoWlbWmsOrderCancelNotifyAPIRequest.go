package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsOrderCancelNotifyAPIRequest
单据取消接口 API请求
taobao.wlb.wms.order.cancel.notify

单据取消接口 */
type TaobaoWlbWmsOrderCancelNotifyAPIRequest struct {
	model.Params
	// 订单类型
	_orderCode string
	// 仓库编码
	_storeCode string
	// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
	_orderType string
}

// New
