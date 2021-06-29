package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提核销订单 API请求
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

// 初始化TaobaoOmniorderStorecollectConsumeRequest对象
func NewTaobaoOmniorderStorecollectConsumeRequest() *TaobaoOmniorderStorecollectConsumeRequest{
    return &TaobaoOmniorderStorecollectConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStorecollectConsumeRequest) GetApiMethodName() string {
    return "taobao.omniorder.storecollect.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStorecollectConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 核销码
func (r *TaobaoOmniorderStorecollectConsumeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoOmniorderStorecollectConsumeRequest) GetCode() string {
    return r.code
}
// MainOrderId Setter
// 淘宝主订单ID
func (r *TaobaoOmniorderStorecollectConsumeRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoOmniorderStorecollectConsumeRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// Operator Setter
// 核销操作人信息
func (r *TaobaoOmniorderStorecollectConsumeRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoOmniorderStorecollectConsumeRequest) GetOperator() string {
    return r.operator
}
