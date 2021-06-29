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
    _nType   string
    // 调用签名
    _strSign   string
    // 参数
    _strContext   string
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
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetNType(_nType string) error {
    r._nType = _nType
    r.Set("n_type", _nType)
    return nil
}

// NType Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetNType() string {
    return r._nType
}
// StrSign Setter
// 调用签名
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrSign(_strSign string) error {
    r._strSign = _strSign
    r.Set("str_sign", _strSign)
    return nil
}

// StrSign Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrSign() string {
    return r._strSign
}
// StrContext Setter
// 参数
func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrContext(_strContext string) error {
    r._strContext = _strContext
    r.Set("str_context", _strContext)
    return nil
}

// StrContext Getter
func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrContext() string {
    return r._strContext
}
