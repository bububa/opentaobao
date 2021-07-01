package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AliExpress退供单明细查询API 
aliexpress.ascp.ro.item.query

AE仓发 单个退供单明细查询
*/
func AliexpressAscpRoItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpRoItemQueryAPIRequest, session string) (*ascpffo.AliexpressAscpRoItemQueryAPIResponse, error) {
    var resp ascpffo.AliexpressAscpRoItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
