package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AliExpress发货单查询API 
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口
*/
func AliexpressAscpFfoQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFfoQueryAPIRequest, session string) (*ascpffo.AliexpressAscpFfoQueryAPIResponse, error) {
    var resp ascpffo.AliexpressAscpFfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
