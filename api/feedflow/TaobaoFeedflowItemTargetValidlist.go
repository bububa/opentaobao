package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemTargetValidlist 获取有权限的定向列表
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
func TaobaoFeedflowItemTargetValidlist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemTargetValidlistAPIRequest, resp *feedflow.TaobaoFeedflowItemTargetValidlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
