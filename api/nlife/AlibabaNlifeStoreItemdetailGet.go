package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
查询商品的详情信息 
alibaba.nlife.store.itemdetail.get

查询零售加平台上单个商品的详情信息
*/
func AlibabaNlifeStoreItemdetailGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreItemdetailGetRequest, session string) (*nlife.AlibabaNlifeStoreItemdetailGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeStoreItemdetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
