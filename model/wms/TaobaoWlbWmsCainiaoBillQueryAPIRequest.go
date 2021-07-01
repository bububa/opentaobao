package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsCainiaoBillQueryAPIRequest
查询单据列表 API请求
taobao.wlb.wms.cainiao.bill.query

查询单据列表 */
type TaobaoWlbWmsCainiaoBillQueryAPIRequest struct {
	model.Params
	// 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
	_startModifiedTime string
	// 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
	_endModifiedTime string
	// 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
	_orderType string
	// 页码。（大于0的整数。默认为1）
	_pageNo int64
	// 每页条数。（每页条数不超过50条。默认为50）
	_pageSize int64
}

// New
