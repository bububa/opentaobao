package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateGet 酒店产品库rate查询
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
func TaobaoXhotelRateGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateGetAPIRequest, resp *xhotelitem.TaobaoXhotelRateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
