package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口订单拉取 
alibaba.wdk.order.list

五道口交易订单拉取接口
*/
func AlibabaWdkOrderList(clt *core.SDKClient, req *wdk.AlibabaWdkOrderListAPIRequest, session string) (*wdk.AlibabaWdkOrderListAPIResponse, error) {
    var resp wdk.AlibabaWdkOrderListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
