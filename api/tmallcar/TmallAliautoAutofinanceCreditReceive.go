package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoAutofinanceCreditReceive 接收授信结果通知
// tmall.aliauto.autofinance.credit.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
func TmallAliautoAutofinanceCreditReceive(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoAutofinanceCreditReceiveAPIRequest, resp *tmallcar.TmallAliautoAutofinanceCreditReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
