package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
驱动保税交易订单发货 
cainiao.bim.tradeorder.consign

驱动保税交易订单发货
*/
func CainiaoBimTradeorderConsign(clt *core.SDKClient, req *wms.CainiaoBimTradeorderConsignRequest, session string) (*wms.CainiaoBimTradeorderConsignResponse, error) {
    var resp wms.CainiaoBimTradeorderConsignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
