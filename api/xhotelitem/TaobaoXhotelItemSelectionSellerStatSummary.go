package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelItemSelectionSellerStatSummary 商家数据-选品整体概况
// taobao.xhotel.item.selection.seller.stat.summary
//
// 商家数据-选品整体概况
func TaobaoXhotelItemSelectionSellerStatSummary(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest, resp *xhotelitem.TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
