package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallServicecenterTpFundsRecoverQuery 服务商资金权益逆向扣回的查询接口
// tmall.servicecenter.tp.funds.recover.query
//
// 服务商资金权益逆向扣回的查询接口
func TmallServicecenterTpFundsRecoverQuery(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallServicecenterTpFundsRecoverQueryAPIRequest, resp *servicecenter.TmallServicecenterTpFundsRecoverQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
