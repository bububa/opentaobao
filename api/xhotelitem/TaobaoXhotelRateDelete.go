package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
rate删除接口 
taobao.xhotel.rate.delete

酒店产品库rate删除
*/
func TaobaoXhotelRateDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateDeleteRequest, session string) (*xhotelitem.TaobaoXhotelRateDeleteAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelRateDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
