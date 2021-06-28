package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
给指定用户创建appkey 
account.aliyuncs.com.CreateApp.2013-07-01

为某个用户创建appkey
*/
func AccountAliyuncsComCreateApp2013-07-01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateApp2013-07-01Request, session string) (*aliyun.AccountAliyuncsComCreateApp2013-07-01APIResponse, error) {
    var resp aliyun.AccountAliyuncsComCreateApp2013-07-01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
