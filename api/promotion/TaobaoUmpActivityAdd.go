package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpActivityAdd 新增优惠活动
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
func TaobaoUmpActivityAdd(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpActivityAddAPIRequest, resp *promotion.TaobaoUmpActivityAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
