package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
财务订单回流 
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据
*/
func AlibabaWdkFinanceOrderBackflow(clt *core.SDKClient, req *wdk.AlibabaWdkFinanceOrderBackflowAPIRequest, session string) (*wdk.AlibabaWdkFinanceOrderBackflowAPIResponse, error) {
    var resp wdk.AlibabaWdkFinanceOrderBackflowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
