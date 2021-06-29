package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送重发码 API请求
taobao.omniorder.dtd.resend

该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
*/
type TaobaoOmniorderDtdResendRequest struct {
    model.Params
    // 淘宝主订单ID
    _mainOrderId   int64
}

// 初始化TaobaoOmniorderDtdResendRequest对象
func NewTaobaoOmniorderDtdResendRequest() *TaobaoOmniorderDtdResendRequest{
    return &TaobaoOmniorderDtdResendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdResendRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.resend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 淘宝主订单ID
func (r *TaobaoOmniorderDtdResendRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniorderDtdResendRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
