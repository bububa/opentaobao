package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
获得当前系统时间 
alibaba.wdk.time.get

获得当前系统时间
*/
func AlibabaWdkTimeGet(clt *core.SDKClient, req *wdk.AlibabaWdkTimeGetRequest, session string) (*wdk.AlibabaWdkTimeGetAPIResponse, error) {
    var resp wdk.AlibabaWdkTimeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
