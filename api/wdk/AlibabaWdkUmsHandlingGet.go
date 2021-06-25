package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
加工单-回流单（新接口） 
alibaba.wdk.ums.handling.get

加工单-回流单（新接口）
*/
func AlibabaWdkUmsHandlingGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsHandlingGetRequest, session string) (*wdk.AlibabaWdkUmsHandlingGetResponse, error) {
    var resp wdk.AlibabaWdkUmsHandlingGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
