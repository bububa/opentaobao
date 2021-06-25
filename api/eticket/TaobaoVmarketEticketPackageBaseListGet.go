package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
根据卖家id，获取关联的所有包 
taobao.vmarket.eticket.package.base.list.get

根据卖家id，获取关联的所有包
*/
func TaobaoVmarketEticketPackageBaseListGet(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketPackageBaseListGetRequest, session string) (*eticket.TaobaoVmarketEticketPackageBaseListGetResponse, error) {
    var resp eticket.TaobaoVmarketEticketPackageBaseListGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
