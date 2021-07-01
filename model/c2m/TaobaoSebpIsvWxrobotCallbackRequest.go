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
type TaobaoSebpIsvWxrobotCallbackAPIRequest struct {
    model.Params
    // 操作类型
    _nType   string
    // 调用签名
    _strSign   string
    // 参数
    _strContext   string
}

// 初始化TaobaoSebpIsvWxrobotCallbackAPIRequest对象
func NewTaobaoSebpIsvWxrobotCallbackRequest() *TaobaoSebpIsvWxrobotCallbackAPIRequest{
    return &TaobaoSebpIsvWxrobotCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvWxrobotCallbackAPIRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.wxrobot.callback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvWxrobotCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NType Setter
// 操作类型
func (r *TaobaoSebpIsvWxrobotCallbackAPIRequest) SetNType(_nType string) error {
    r._nType = _nType
    r.Set("n_type", _nType)
    return nil
}

// NType Getter
func (r TaobaoSebpIsvWxrobotCallbackAPIRequest) GetNType() string {
    return r._nType
}
// StrSign Setter
// 调用签名
func (r *TaobaoSebpIsvWxrobotCallbackAPIRequest) SetStrSign(_strSign string) error {
    r._strSign = _strSign
    r.Set("str_sign", _strSign)
    return nil
}

// StrSign Getter
func (r TaobaoSebpIsvWxrobotCallbackAPIRequest) GetStrSign() string {
    return r._strSign
}
// StrContext Setter
// 参数
func (r *TaobaoSebpIsvWxrobotCallbackAPIRequest) SetStrContext(_strContext string) error {
    r._strContext = _strContext
    r.Set("str_context", _strContext)
    return nil
}

// StrContext Getter
func (r TaobaoSebpIsvWxrobotCallbackAPIRequest) GetStrContext() string {
    return r._strContext
}
