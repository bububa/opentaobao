package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdzonePage 信息流单元下查看绑定资源位
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
func TaobaoFeedflowItemAdgroupAdzonePage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzonePageAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAdzonePageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
