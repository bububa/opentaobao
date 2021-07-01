package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
混淆nick转openid 
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
*/
func TaobaoTopOpenidConvert(clt *core.SDKClient, req *util.TaobaoTopOpenidConvertAPIRequest, session string) (*util.TaobaoTopOpenidConvertAPIResponse, error) {
    var resp util.TaobaoTopOpenidConvertAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
