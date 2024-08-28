package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelAdd 酒店新增接口（ID重复自动更新）
// taobao.xhotel.add
//
// 添加酒店或更新酒店
func TaobaoXhotelAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelAddAPIRequest, resp *xhotelitem.TaobaoXhotelAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
