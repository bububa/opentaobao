package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdzoneList 批量查询可用广告位列表
// taobao.feedflow.item.adzone.list
//
// 批量查询可用广告位列表
func TaobaoFeedflowItemAdzoneList(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdzoneListAPIRequest, resp *feedflow.TaobaoFeedflowItemAdzoneListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
