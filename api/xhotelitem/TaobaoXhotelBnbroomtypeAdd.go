package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbroomtypeAdd 民宿新增房源
// taobao.xhotel.bnbroomtype.add
//
// 添加民宿房源
func TaobaoXhotelBnbroomtypeAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbroomtypeAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbroomtypeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
