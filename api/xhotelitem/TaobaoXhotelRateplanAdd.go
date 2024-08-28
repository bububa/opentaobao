package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanAdd 酒店产品库rateplan添加
// taobao.xhotel.rateplan.add
//
// 酒店产品库rateplan
func TaobaoXhotelRateplanAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanAddAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
