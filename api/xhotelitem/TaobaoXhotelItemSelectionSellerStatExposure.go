package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelItemSelectionSellerStatExposure 选品曝光趋势
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
func TaobaoXhotelItemSelectionSellerStatExposure(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIRequest, resp *xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
