package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanGet 价格计划rateplan查询
// taobao.xhotel.rateplan.get
//
// 酒店产品库rateplan查询
func TaobaoXhotelRateplanGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanGetAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
