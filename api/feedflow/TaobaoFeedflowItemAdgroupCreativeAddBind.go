package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupCreativeAddBind 信息流新增并且绑定创意
// taobao.feedflow.item.adgroup.creative.add.bind
//
// 信息流新增并且绑定创意
func TaobaoFeedflowItemAdgroupCreativeAddBind(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
