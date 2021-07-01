package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTpFundsRecoverQueryAPIRequest
服务商资金权益逆向扣回的查询接口 API请求
tmall.servicecenter.tp.funds.recover.query

服务商资金权益逆向扣回的查询接口 */
type TmallServicecenterTpFundsRecoverQueryAPIRequest struct {
	model.Params
	// query入参
	_query *TpFundsRecoverQuery
}

// New
