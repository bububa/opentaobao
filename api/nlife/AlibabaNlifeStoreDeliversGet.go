package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
获取门店采购单下的发货单列表 
alibaba.nlife.store.delivers.get

获取门店采购单下的发货单列表
*/
func AlibabaNlifeStoreDeliversGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreDeliversGetRequest, session string) (*nlife.AlibabaNlifeStoreDeliversGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeStoreDeliversGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
