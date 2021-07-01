package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreDeliversGetAPIRequest
获取门店采购单下的发货单列表 API请求
alibaba.nlife.store.delivers.get

获取门店采购单下的发货单列表 */
type AlibabaNlifeStoreDeliversGetAPIRequest struct {
	model.Params
	// 门店采购订单号
	_tradeNo string
	// 零售商的门店id
	_storeId int64
	// 每页的数量
	_pageSize int64
	// 查询的页码
	_pageNo int64
}

// New
