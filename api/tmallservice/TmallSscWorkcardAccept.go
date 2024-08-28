package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallSscWorkcardAccept 服务商接单完成
// tmall.ssc.workcard.accept
//
// 工单完结
func TmallSscWorkcardAccept(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallSscWorkcardAcceptAPIRequest, resp *tmallservice.TmallSscWorkcardAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
