package miniappopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappopen"
)

/* 
生成淘宝小程序短链接 
taobao.miniapp.shorturl.create

提供淘宝小程序短链接生成的能力，只允许对淘宝小程序对应的域名：https://m.duanqu.com/ 生成对应的短链接，其他域名无效
【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】
*/
func TaobaoMiniappShorturlCreate(clt *core.SDKClient, req *miniappopen.TaobaoMiniappShorturlCreateRequest, session string) (*miniappopen.TaobaoMiniappShorturlCreateAPIResponse, error) {
    var resp miniappopen.TaobaoMiniappShorturlCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
