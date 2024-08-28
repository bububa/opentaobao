package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelMultiplerateDelete 复杂价格删除
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
func TaobaoXhotelMultiplerateDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelMultiplerateDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
