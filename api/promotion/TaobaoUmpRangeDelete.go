package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpRangeDelete 删除活动范围
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
func TaobaoUmpRangeDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpRangeDeleteAPIRequest, resp *promotion.TaobaoUmpRangeDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
