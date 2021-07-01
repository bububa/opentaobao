package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AliExpress销退单明细查询API 
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API
*/
func AliexpressAscpFroItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFroItemQueryAPIRequest, session string) (*ascpffo.AliexpressAscpFroItemQueryAPIResponse, error) {
    var resp ascpffo.AliexpressAscpFroItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
