package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生成淘宝小程序短链接 APIRequest
taobao.miniapp.shorturl.create

提供淘宝小程序短链接生成的能力，只允许对淘宝小程序对应的域名：https://m.duanqu.com/ 生成对应的短链接，其他域名无效
【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】
*/
type TaobaoMiniappShorturlCreateRequest struct {
    model.Params

    // 小程序链接地址。说明：链接地址，只允许https协议，域名只支持m.duanqu.com，链接必须包含_ariver_appid参数，链接不能够包含spm、short_name、app、tb_url_time_stamp这些系统保留参数
    miniappUrl   string 

}

func NewTaobaoMiniappShorturlCreateRequest() *TaobaoMiniappShorturlCreateRequest{
    return &TaobaoMiniappShorturlCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappShorturlCreateRequest) GetApiMethodName() string {
    return "taobao.miniapp.shorturl.create"
}

func (r TaobaoMiniappShorturlCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappShorturlCreateRequest) SetMiniappUrl(miniappUrl string) error {
    r.miniappUrl = miniappUrl
    r.Set("miniapp_url", miniappUrl)
    return nil
}

func (r TaobaoMiniappShorturlCreateRequest) GetMiniappUrl() string {
    return r.miniappUrl
}

