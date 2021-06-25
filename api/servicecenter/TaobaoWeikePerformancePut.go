package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
提交客服绩效接口 
taobao.weike.performance.put

提交客服绩效接口
*/
func TaobaoWeikePerformancePut(clt *core.SDKClient, req *servicecenter.TaobaoWeikePerformancePutRequest, session string) (*servicecenter.TaobaoWeikePerformancePutResponse, error) {
    var resp servicecenter.TaobaoWeikePerformancePutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
