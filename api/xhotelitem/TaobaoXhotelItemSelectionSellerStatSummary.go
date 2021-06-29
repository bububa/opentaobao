package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
商家数据-选品整体概况 
taobao.xhotel.item.selection.seller.stat.summary

商家数据-选品整体概况
*/
func TaobaoXhotelItemSelectionSellerStatSummary(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatSummaryRequest, session string) (*xhotelitem.TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
