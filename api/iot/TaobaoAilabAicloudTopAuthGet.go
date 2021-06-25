package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
登陆 
taobao.ailab.aicloud.top.auth.get

登陆
*/
func TaobaoAilabAicloudTopAuthGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopAuthGetRequest, session string) (*iot.TaobaoAilabAicloudTopAuthGetResponse, error) {
    var resp iot.TaobaoAilabAicloudTopAuthGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
