package alime

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alime"
)

/* 
获取用户免登录令牌 
taobao.alime.user.token.get

根据第三账号信息获取用户的免登录令牌
*/
func TaobaoAlimeUserTokenGet(clt *core.SDKClient, req *alime.TaobaoAlimeUserTokenGetAPIRequest, session string) (*alime.TaobaoAlimeUserTokenGetAPIResponse, error) {
    var resp alime.TaobaoAlimeUserTokenGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
