package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelitemselectionsellerstatsummary 商家数据-选品整体概况
// taobao.xhotel.item.selection.seller.stat.summary
//
// 商家数据-选品整体概况
func Taobaoxhotelitemselectionsellerstatsummary(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelitemselectionsellerstatsummaryAPIRequest, session string) (*xhotelitem.TaobaoxhotelitemselectionsellerstatsummaryAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelitemselectionsellerstatsummaryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
