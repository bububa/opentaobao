package alisports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alisports"
)

/* 
获取会员信息 
alibaba.alisports.passport.account.getaccountinfo

获取阿里体育会员信息
*/
func AlibabaAlisportsPassportAccountGetaccountinfo(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIResponse, error) {
    var resp alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
