package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
创建阿里云账号 
account.aliyuncs.com.CreateAliyunAccount.2013-07-01

根据给定的阿里云账号，密码以及手机号创建阿里云账号
*/
func AccountAliyuncsComCreateAliyunAccount2013_07_01(clt *core.SDKClient, req *aliyun.AccountAliyuncsComCreateAliyunAccount2013_07_01Request, session string) (*aliyun.AccountAliyuncsComCreateAliyunAccount2013_07_01APIResponse, error) {
    var resp aliyun.AccountAliyuncsComCreateAliyunAccount2013_07_01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
