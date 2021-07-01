package xhoteloffline

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhoteloffline"
)

/* 
线下信用住结账专用接口 
taobao.xhotel.order.offline.settle.put

线下信用住结账专用接口
*/
func TaobaoXhotelOrderOfflineSettlePut(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderOfflineSettlePutAPIRequest, session string) (*xhoteloffline.TaobaoXhotelOrderOfflineSettlePutAPIResponse, error) {
    var resp xhoteloffline.TaobaoXhotelOrderOfflineSettlePutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
