package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无交易类凭证核销 APIRequest
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

func NewTaobaoVmarketEticketFlowConsumeRequest() *TaobaoVmarketEticketFlowConsumeRequest{
    return &TaobaoVmarketEticketFlowConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.flow.consume"
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoVmarketEticketFlowConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetCode() string {
    return r.code
}

func (r *TaobaoVmarketEticketFlowConsumeRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoVmarketEticketFlowConsumeRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoVmarketEticketFlowConsumeRequest) GetOperator() string {
    return r.operator
}

