package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
获取企业下的采购单列表 
alibaba.nlife.b2b.trade.list

获取指定门店下的采购单列表
*/
func AlibabaNlifeB2bTradeList(clt *core.SDKClient, req *nlife.AlibabaNlifeB2bTradeListRequest, session string) (*nlife.AlibabaNlifeB2bTradeListAPIResponse, error) {
    var resp nlife.AlibabaNlifeB2bTradeListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
