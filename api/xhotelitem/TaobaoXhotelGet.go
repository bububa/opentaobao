package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelGet 酒店查询接口
// taobao.xhotel.get
//
// 酒店查询接口
func TaobaoXhotelGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetAPIRequest, resp *xhotelitem.TaobaoXhotelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
