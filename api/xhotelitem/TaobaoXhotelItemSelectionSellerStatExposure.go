package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelItemSelectionSellerStatExposure 选品曝光趋势
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
func TaobaoXhotelItemSelectionSellerStatExposure(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIRequest, resp *xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
