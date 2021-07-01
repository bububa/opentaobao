package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTpFundsSendQueryAPIRequest
服务商资金权益发放的查询接口 API请求
tmall.servicecenter.tp.funds.send.query

服务商资金权益发放结果的查询接口 */
type TmallServicecenterTpFundsSendQueryAPIRequest struct {
	model.Params
	// 入参对象
	_query *TpFundsSendQuery
}

// New
