package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
盒马帮采购确认订单接口 
alibaba.wdk.wholesale.order.commit

盒马帮采购确认订单接口
*/
func AlibabaWdkWholesaleOrderCommit(clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleOrderCommitRequest, session string) (*wdk.AlibabaWdkWholesaleOrderCommitAPIResponse, error) {
    var resp wdk.AlibabaWdkWholesaleOrderCommitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
