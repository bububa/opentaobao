package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取TOP通道解密秘钥 
taobao.top.secret.get

top sdk通过api获取对应解密秘钥
*/
func TaobaoTopSecretGet(clt *core.SDKClient, req *util.TaobaoTopSecretGetRequest, session string) (*util.TaobaoTopSecretGetAPIResponse, error) {
    var resp util.TaobaoTopSecretGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
