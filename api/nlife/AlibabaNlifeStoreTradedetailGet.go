package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
查询采购单详情信息 
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息
*/
func AlibabaNlifeStoreTradedetailGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreTradedetailGetRequest, session string) (*nlife.AlibabaNlifeStoreTradedetailGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeStoreTradedetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
