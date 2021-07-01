package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
查询发货单详情 
alibaba.nlife.store.deliverdetail.get

查询发货单详情
*/
func AlibabaNlifeStoreDeliverdetailGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreDeliverdetailGetAPIRequest, session string) (*nlife.AlibabaNlifeStoreDeliverdetailGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeStoreDeliverdetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
