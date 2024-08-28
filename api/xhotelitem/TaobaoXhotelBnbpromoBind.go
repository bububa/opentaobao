package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoBind 自促活动绑定接口
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
func TaobaoXhotelBnbpromoBind(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoBindAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
