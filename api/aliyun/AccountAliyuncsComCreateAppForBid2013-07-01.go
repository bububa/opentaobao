package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
给当前渠道下的用户创建app 
account.aliyuncs.com.CreateAppForBid.2013-07-01

给自己渠道下的用户创建app
*/
func AccountAliyuncsComCreateAppForBid2013-07-01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAppForBid2013-07-01Request, session string) (*aliyun.AccountAliyuncsComCreateAppForBid2013-07-01Response, error) {
    var resp aliyun.AccountAliyuncsComCreateAppForBid2013-07-01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
