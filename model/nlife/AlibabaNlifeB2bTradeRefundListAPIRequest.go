package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2bTradeRefundListAPIRequest
获取采购退货单列表 API请求
alibaba.nlife.b2b.trade.refund.list

获取采购退货单列表 */
type AlibabaNlifeB2bTradeRefundListAPIRequest struct {
	model.Params
	// 采购退货单创建时间开始范围
	_startEffectiveDate string
	// 采购退货单创建时间结束范围
	_endEffectiveDate string
	// 查询的页数
	_pageNo int64
	// 每页的数量
	_pageSize int64
	// 企业Id
	_entId int64
}

// New
