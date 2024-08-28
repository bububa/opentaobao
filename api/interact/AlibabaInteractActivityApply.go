package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractActivityApply ISV报名官方活动(中心化流量)
// alibaba.interact.activity.apply
//
// 支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
func AlibabaInteractActivityApply(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractActivityApplyAPIRequest, resp *interact.AlibabaInteractActivityApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
