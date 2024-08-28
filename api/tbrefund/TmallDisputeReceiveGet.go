package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TmallDisputeReceiveGet 天猫逆向纠纷查询
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
func TmallDisputeReceiveGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TmallDisputeReceiveGetAPIRequest, resp *tbrefund.TmallDisputeReceiveGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
