package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
新增专享房价 
taobao.xhotel.rate.add

酒店产品库rate添加
*/
func TaobaoXhotelRateAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateAddAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelRateAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
