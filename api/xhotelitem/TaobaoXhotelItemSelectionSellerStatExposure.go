package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
选品曝光趋势 
taobao.xhotel.item.selection.seller.stat.exposure

用于提供给商家获取选品曝光趋势
*/
func TaobaoXhotelItemSelectionSellerStatExposure(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIRequest, session string) (*xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelItemSelectionSellerStatExposureAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
