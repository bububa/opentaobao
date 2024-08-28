package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanUpdate 价格计划rateplan更新或添加
// taobao.xhotel.rateplan.update
//
// 酒店产品库rateplan更新或添加
func TaobaoXhotelRateplanUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
