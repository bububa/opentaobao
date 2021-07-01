package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
批量消息确认 
alibaba.wdk.ums.retrieve.batch.confirm

批量消息确认
*/
func AlibabaWdkUmsRetrieveBatchConfirm(clt *core.SDKClient, req *wdk.AlibabaWdkUmsRetrieveBatchConfirmAPIRequest, session string) (*wdk.AlibabaWdkUmsRetrieveBatchConfirmAPIResponse, error) {
    var resp wdk.AlibabaWdkUmsRetrieveBatchConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
