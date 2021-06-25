package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
为bid用户创建账号 
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记
*/
func AccountAliyuncsComCreateAliyunAccountForBid2013-07-01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Request, session string) (*aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013-07-01Response, error) {
    var resp aliyun.AccountAliyuncsComCreateAliyunAccountForBid2013-07-01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
