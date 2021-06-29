package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv机器人回调接口 API请求
taobao.sebp.isv.wxrobot.callback

机器人入群回调，进行校验、功能开通等操作
*/
type TaobaoSebpIsvWxrobotCallbackRequest struct {
    model.Params
    // 操作类型
    nType   string
    // 调用签名
    strSign   string
    // 参数
    strContext   string
}

// 初始化TaobaoSebpIsvWxrobotCallbackRequest对象
func NewTaobaoSebpIsvWxrobotCallbackRequest() *TaobaoSebpIsvWxrobotCallbackRequest{
    return &TaobaoSebpIsvWxrobotCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.wxrobot.callback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NType Setter
// 操作类型
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetNType(nType string) error {
    r.nType = nType
    r.Set("n_type", nType)
    return nil
}

// NType Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetNType() string {
    return r.nType
}
// StrSign Setter
// 调用签名
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrSign(strSign string) error {
    r.strSign = strSign
    r.Set("str_sign", strSign)
    return nil
}

// StrSign Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrSign() string {
    return r.strSign
}
// StrContext Setter
// 参数
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrContext(strContext string) error {
    r.strContext = strContext
    r.Set("str_context", strContext)
    return nil
}

// StrContext Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrContext() string {
    return r.strContext
}
