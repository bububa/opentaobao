package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPerfectPerformanceItemQueryAPIRequest
商品完美履约信息查询 API请求
alibaba.perfect.performance.item.query

同城零售商品完美履约信息查询 */
type AlibabaPerfectPerformanceItemQueryAPIRequest struct {
	model.Params
	// 查询入参
	_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq
}

// New
