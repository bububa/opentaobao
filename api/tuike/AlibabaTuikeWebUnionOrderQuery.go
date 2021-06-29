package tuike

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tuike"
)

/* 
推客网盟合作抽佣订单查询接口 
alibaba.tuike.web.union.order.query

推客网盟合作抽佣订单查询接口
*/
func AlibabaTuikeWebUnionOrderQuery(clt *core.SDKClient, req *tuike.AlibabaTuikeWebUnionOrderQueryRequest, session string) (*tuike.AlibabaTuikeWebUnionOrderQueryAPIResponse, error) {
    var resp tuike.AlibabaTuikeWebUnionOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}