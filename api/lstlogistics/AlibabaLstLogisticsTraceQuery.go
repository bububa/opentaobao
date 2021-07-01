package lstlogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstlogistics"
)

/* 
供应商-异云-查询运单物流追踪信息 
alibaba.lst.logistics.trace.query

查询LP单物流追踪信息
*/
func AlibabaLstLogisticsTraceQuery(clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsTraceQueryAPIRequest, session string) (*lstlogistics.AlibabaLstLogisticsTraceQueryAPIResponse, error) {
    var resp lstlogistics.AlibabaLstLogisticsTraceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
