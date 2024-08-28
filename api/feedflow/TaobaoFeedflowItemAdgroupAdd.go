package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdd 信息流增加单元
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
func TaobaoFeedflowItemAdgroupAdd(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAddAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
