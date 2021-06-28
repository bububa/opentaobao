package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口查询同步订单 
alibaba.wdk.syncedorder.query

外部商户查询同步到五道口的订单
*/
func AlibabaWdkSyncedorderQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSyncedorderQueryRequest, session string) (*wdk.AlibabaWdkSyncedorderQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkSyncedorderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
