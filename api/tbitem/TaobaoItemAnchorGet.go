package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemAnchorGet 获取可用宝贝描述规范化模块
// taobao.item.anchor.get
//
// 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
func TaobaoItemAnchorGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemAnchorGetAPIRequest, resp *tbitem.TaobaoItemAnchorGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
