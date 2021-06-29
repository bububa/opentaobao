package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv机器人回调接口 APIRequest
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

func NewTaobaoSebpIsvWxrobotCallbackRequest() *TaobaoSebpIsvWxrobotCallbackRequest{
    return &TaobaoSebpIsvWxrobotCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSebpIsvWxrobotCallbackRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.wxrobot.callback"
}

func (r TaobaoSebpIsvWxrobotCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetNType(nType string) error {
    r.nType = nType
    r.Set("n_type", nType)
    return nil
}

func (r TaobaoSebpIsvWxrobotCallbackRequest) GetNType() string {
    return r.nType
}

func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrSign(strSign string) error {
    r.strSign = strSign
    r.Set("str_sign", strSign)
    return nil
}

func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrSign() string {
    return r.strSign
}

func (r *TaobaoSebpIsvWxrobotCallbackRequest) SetStrContext(strContext string) error {
    r.strContext = strContext
    r.Set("str_context", strContext)
    return nil
}

func (r TaobaoSebpIsvWxrobotCallbackRequest) GetStrContext() string {
    return r.strContext
}

