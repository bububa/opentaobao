package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2bTradeListAPIRequest
获取企业下的采购单列表 API请求
alibaba.nlife.b2b.trade.list

获取指定门店下的采购单列表 */
type AlibabaNlifeB2bTradeListAPIRequest struct {
	model.Params
	// 采购单生效时间开始范围
	_startEffectiveDate string
	// 采购单生效时间结束范围
	_endEffectiveDate string
	// 查询的页码
	_pageNo int64
	// 每页的数量
	_pageSize int64
	// 企业ID
	_entId int64
}

// New
