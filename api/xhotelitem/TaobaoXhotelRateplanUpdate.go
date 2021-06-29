package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
价格计划rateplan更新或添加 
taobao.xhotel.rateplan.update

酒店产品库rateplan更新或添加
*/
func TaobaoXhotelRateplanUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanUpdateRequest, session string) (*xhotelitem.TaobaoXhotelRateplanUpdateAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelRateplanUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
