package yunosminiapp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/yunosminiapp"
)

/* 
调用活动接口 
yunos.miniapp.activity.call

用于小程序调用活动接口
*/
func YunosMiniappActivityCall(clt *core.SDKClient, req *yunosminiapp.YunosMiniappActivityCallRequest, session string) (*yunosminiapp.YunosMiniappActivityCallAPIResponse, error) {
    var resp yunosminiapp.YunosMiniappActivityCallAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
