package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
通过物流订单编号查询物流信息 
taobao.wlb.tmsorder.query

通过物流订单编号分页查询物流信息
*/
func TaobaoWlbTmsorderQuery(clt *core.SDKClient, req *wlb.TaobaoWlbTmsorderQueryRequest, session string) (*wlb.TaobaoWlbTmsorderQueryAPIResponse, error) {
    var resp wlb.TaobaoWlbTmsorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
