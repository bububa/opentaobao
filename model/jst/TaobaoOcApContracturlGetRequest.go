package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按用户获取支付宝代扣协议链接地址 API请求
taobao.oc.ap.contracturl.get

按用户获取支付宝代扣协议链接地址
*/
type TaobaoOcApContracturlGetRequest struct {
    model.Params
}

// 初始化TaobaoOcApContracturlGetRequest对象
func NewTaobaoOcApContracturlGetRequest() *TaobaoOcApContracturlGetRequest{
    return &TaobaoOcApContracturlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcApContracturlGetRequest) GetApiMethodName() string {
    return "taobao.oc.ap.contracturl.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcApContracturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
