package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbreviewAdd 对外开放评论接口
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
func TaobaoXhotelBnbreviewAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbreviewAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbreviewAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
