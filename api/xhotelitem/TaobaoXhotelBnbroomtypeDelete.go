package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbroomtypeDelete 民宿房源删除接口
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
func TaobaoXhotelBnbroomtypeDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
