package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelXitemQuery 查询 x 元素
// taobao.xhotel.xitem.query
//
// 查询 x 元素
func TaobaoXhotelXitemQuery(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelXitemQueryAPIRequest, resp *xhotelitem.TaobaoXhotelXitemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
