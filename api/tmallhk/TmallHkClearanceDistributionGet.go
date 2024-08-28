package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceDistributionGet 分销供应商获取清关材料
// tmall.hk.clearance.distribution.get
//
// 供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
func TmallHkClearanceDistributionGet(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallHkClearanceDistributionGetAPIRequest, resp *tmallhk.TmallHkClearanceDistributionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
