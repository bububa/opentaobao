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
    outerId   string
    // 凭证码
    code   string
    // 淘宝业务提供的业务类型值，请联系相关业务运营取得
    bizType   int64
    // 核销操作人
    operator   string
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
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetOuterId() string {
    return r.outerId
}
// Code Setter
// 凭证码
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetCode() string {
    return r.code
}
// BizType Setter
// 淘宝业务提供的业务类型值，请联系相关业务运营取得
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetBizType() int64 {
    return r.bizType
}
// Operator Setter
// 核销操作人
func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoVmarketEticketFlowConsumeRequest) GetOperator() string {
    return r.operator
}
