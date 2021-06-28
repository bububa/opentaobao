package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
获取电子凭证预约门店信息 
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息
*/
func TaobaoVmarketEticketStoreGet(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketStoreGetRequest, session string) (*eticket.TaobaoVmarketEticketStoreGetAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketStoreGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
