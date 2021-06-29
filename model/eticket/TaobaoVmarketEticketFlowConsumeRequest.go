package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无交易类凭证核销 API请求
taobao.vmarket.eticket.flow.consume

无交易类凭证核销
*/
type TaobaoVmarketEticketFlowConsumeRequest struct {
    model.Params
    // 业务单号
    _outerId   string
    // 凭证码
    _code   string
    // 淘宝业务提供的业务类型值，请联系相关业务运营取得
    _bizType   int64
    // 核销操作人
    _operator   string
}

// 初始化TaobaoVmarketEticketFlowConsumeRequest对象
func NewTaobaoVmarketEticketFlowConsumeRequest() *TaobaoVmarketEticketFlowConsumeRequest{
    return &TaobaoVmarketEticketFlowConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFlowConsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.flow.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFlowConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 业务单号
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetOuterId() string {
    return r._outerId
}
// Code Setter
// 凭证码
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetCode() string {
    return r._code
}
// BizType Setter
// 淘宝业务提供的业务类型值，请联系相关业务运营取得
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetBizType() int64 {
    return r._bizType
}
// Operator Setter
// 核销操作人
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetOperator() string {
    return r._operator
}
