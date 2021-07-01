package miniapp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniapp"
)

/* 
外部触发云函数 
taobao.miniapp.cloud.function.invoke

用户isv从外部触发聚石塔云函数的执行。
*/
func TaobaoMiniappCloudFunctionInvoke(clt *core.SDKClient, req *miniapp.TaobaoMiniappCloudFunctionInvokeAPIRequest, session string) (*miniapp.TaobaoMiniappCloudFunctionInvokeAPIResponse, error) {
    var resp miniapp.TaobaoMiniappCloudFunctionInvokeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
