package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupPage 查询单元列表
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
func TaobaoFeedflowItemAdgroupPage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupPageAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
