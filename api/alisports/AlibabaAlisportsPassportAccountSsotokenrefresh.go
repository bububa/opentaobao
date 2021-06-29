package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
sso_token刷新 
alibaba.alisports.passport.account.ssotokenrefresh

sso_token刷新
*/
func AlibabaAlisportsPassportAccountSsotokenrefresh(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountSsotokenrefreshRequest, session string) (*alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse, error) {
    var resp alisports.AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
