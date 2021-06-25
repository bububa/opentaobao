package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
系列品变更-更新系列 
alibaba.wdk.series.edit

系列品变更-更新系列
*/
func AlibabaWdkSeriesEdit(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesEditRequest, session string) (*wdk.AlibabaWdkSeriesEditResponse, error) {
    var resp wdk.AlibabaWdkSeriesEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
