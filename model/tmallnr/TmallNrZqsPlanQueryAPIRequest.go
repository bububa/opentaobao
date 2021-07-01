package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrZqsPlanQueryAPIRequest
周期送配送明细查询 API请求
tmall.nr.zqs.plan.query

周期送配送明细查询 */
type TmallNrZqsPlanQueryAPIRequest struct {
	model.Params
	// 交易子订单id
	_detailOrderId int64
}

// New
