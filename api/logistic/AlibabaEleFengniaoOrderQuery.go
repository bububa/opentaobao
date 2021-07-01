package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
查询订单基本信息 
alibaba.ele.fengniao.order.query

查询订单基本信息
*/
func AlibabaEleFengniaoOrderQuery(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoOrderQueryAPIRequest, session string) (*logistic.AlibabaEleFengniaoOrderQueryAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
