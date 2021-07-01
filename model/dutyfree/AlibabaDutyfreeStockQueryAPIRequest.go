package dutyfree

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDutyfreeStockQueryAPIRequest
对外库存查询接口 API请求
alibaba.dutyfree.stock.query

对外部服务提供库存查询接口 */
type AlibabaDutyfreeStockQueryAPIRequest struct {
	model.Params
	// 条形码
	_barCode string
}

// New
