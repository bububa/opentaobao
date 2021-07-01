package xhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotel"
)

/* 
卖家服务指数查询 
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询
*/
func TaobaoXhotelDataServiceSellerServiceindex(clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceSellerServiceindexAPIRequest, session string) (*xhotel.TaobaoXhotelDataServiceSellerServiceindexAPIResponse, error) {
    var resp xhotel.TaobaoXhotelDataServiceSellerServiceindexAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
