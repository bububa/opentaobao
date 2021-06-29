package perfect

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/perfect"
)

/* 
商品完美履约信息查询 
alibaba.perfect.performance.item.query

同城零售商品完美履约信息查询
*/
func AlibabaPerfectPerformanceItemQuery(clt *core.SDKClient, req *perfect.AlibabaPerfectPerformanceItemQueryRequest, session string) (*perfect.AlibabaPerfectPerformanceItemQueryAPIResponse, error) {
    var resp perfect.AlibabaPerfectPerformanceItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
