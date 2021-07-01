package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
贩卖机扫码查询领奖状态 
alibaba.alihealth.codeseller.getuseraward

贩卖机扫码查询领奖状态
*/
func AlibabaAlihealthCodesellerGetuseraward(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthCodesellerGetuserawardAPIRequest, session string) (*drugtrace.AlibabaAlihealthCodesellerGetuserawardAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthCodesellerGetuserawardAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
