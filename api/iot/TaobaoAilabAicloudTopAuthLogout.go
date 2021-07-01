package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
登出 
taobao.ailab.aicloud.top.auth.logout

登出
*/
func TaobaoAilabAicloudTopAuthLogout(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopAuthLogoutAPIRequest, session string) (*iot.TaobaoAilabAicloudTopAuthLogoutAPIResponse, error) {
    var resp iot.TaobaoAilabAicloudTopAuthLogoutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
