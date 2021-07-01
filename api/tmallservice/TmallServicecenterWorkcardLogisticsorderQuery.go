package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
物流单信息查询 
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
func TmallServicecenterWorkcardLogisticsorderQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardLogisticsorderQueryAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardLogisticsorderQueryAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardLogisticsorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
