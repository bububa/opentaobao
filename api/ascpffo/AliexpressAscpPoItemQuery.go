package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AliExpress采购单明细查询API 
aliexpress.ascp.po.item.query

AE 供应链仓发 采购单明细查询
*/
func AliexpressAscpPoItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpPoItemQueryAPIRequest, session string) (*ascpffo.AliexpressAscpPoItemQueryAPIResponse, error) {
    var resp ascpffo.AliexpressAscpPoItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
