package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
运营商删除用户的appkey 
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。
*/
func AccountAliyuncsComDeleteAppForBid2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest, session string) (*aliyun.AccountAliyuncsComDeleteAppForBid2013_07_01APIResponse, error) {
    var resp aliyun.AccountAliyuncsComDeleteAppForBid2013_07_01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
