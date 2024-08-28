package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallServicecenterTpFundsSendQuery 服务商资金权益发放的查询接口
// tmall.servicecenter.tp.funds.send.query
//
// 服务商资金权益发放结果的查询接口
func TmallServicecenterTpFundsSendQuery(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallServicecenterTpFundsSendQueryAPIRequest, resp *servicecenter.TmallServicecenterTpFundsSendQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
