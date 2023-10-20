package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallServicecenterTpFundsRecoverQuery 服务商资金权益逆向扣回的查询接口
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
func TmallServicecenterTpFundsRecoverQuery(clt *core.SDKClient, req *servicecenter.TmallServicecenterTpFundsRecoverQueryAPIRequest, resp *servicecenter.TmallServicecenterTpFundsRecoverQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
