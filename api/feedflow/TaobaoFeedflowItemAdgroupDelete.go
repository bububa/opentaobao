package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupDelete 根据单元id删除单元
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
func TaobaoFeedflowItemAdgroupDelete(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
