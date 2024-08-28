package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemPropimgDelete 删除属性图片
// taobao.item.propimg.delete
//
// 删除propimg_id 所指定的商品属性图片 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;propimg_id对应的属性图片需要属于num_iid对应的商品
func TaobaoItemPropimgDelete(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemPropimgDeleteAPIRequest, resp *tbitem.TaobaoItemPropimgDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
