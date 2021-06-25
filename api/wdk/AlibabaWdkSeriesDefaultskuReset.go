package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
系列品商品变更-重置默认商品 
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品
*/
func AlibabaWdkSeriesDefaultskuReset(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesDefaultskuResetRequest, session string) (*wdk.AlibabaWdkSeriesDefaultskuResetResponse, error) {
    var resp wdk.AlibabaWdkSeriesDefaultskuResetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
