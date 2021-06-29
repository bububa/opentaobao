package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
登陆用户 
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。
*/
func AlibabaAlifanyiMarketLogin(clt *core.SDKClient, req *seaking.AlibabaAlifanyiMarketLoginRequest, session string) (*seaking.AlibabaAlifanyiMarketLoginAPIResponse, error) {
    var resp seaking.AlibabaAlifanyiMarketLoginAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
