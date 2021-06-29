package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AliExpress采购单查询API 
aliexpress.ascp.po.query

AE仓发业务采购单查询
*/
func AliexpressAscpPoQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpPoQueryRequest, session string) (*ascpffo.AliexpressAscpPoQueryAPIResponse, error) {
    var resp ascpffo.AliexpressAscpPoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
