package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoGet 民宿查询营销活动
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
func TaobaoXhotelBnbpromoGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoGetAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
