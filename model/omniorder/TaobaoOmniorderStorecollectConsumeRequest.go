package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提核销订单 APIRequest
taobao.omniorder.storecollect.consume

全渠道门店自提核销订单
*/
type TaobaoOmniorderStorecollectConsumeRequest struct {
    model.Params

    // 核销码
    code   string 

    // 淘宝主订单ID
    mainOrderId   int64 

    // 核销操作人信息
    operator   string 

}

func NewTaobaoOmniorderStorecollectConsumeRequest() *TaobaoOmniorderStorecollectConsumeRequest{
    return &TaobaoOmniorderStorecollectConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStorecollectConsumeRequest) GetApiMethodName() string {
    return "taobao.omniorder.storecollect.consume"
}

func (r TaobaoOmniorderStorecollectConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStorecollectConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoOmniorderStorecollectConsumeRequest) GetCode() string {
    return r.code
}

func (r *TaobaoOmniorderStorecollectConsumeRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoOmniorderStorecollectConsumeRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoOmniorderStorecollectConsumeRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoOmniorderStorecollectConsumeRequest) GetOperator() string {
    return r.operator
}

