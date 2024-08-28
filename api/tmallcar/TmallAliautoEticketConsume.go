package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoEticketConsume 天猫汽车二轮车电子凭证核销
// tmall.aliauto.eticket.consume
//
// 天猫汽车二轮车行业门店电子凭证核销
func TmallAliautoEticketConsume(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoEticketConsumeAPIRequest, resp *tmallcar.TmallAliautoEticketConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
