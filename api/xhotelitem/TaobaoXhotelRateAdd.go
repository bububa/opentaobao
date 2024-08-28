package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateAdd 新增专享房价
// taobao.xhotel.rate.add
//
// 酒店产品库rate添加
func TaobaoXhotelRateAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateAddAPIRequest, resp *xhotelitem.TaobaoXhotelRateAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
