package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款信息审核 API请求
taobao.rdc.aligenius.refunds.check

根据退款信息，对退款单进行审核
*/
type TaobaoRdcAligeniusRefundsCheckRequest struct {
    model.Params
    // 入参
    _param   *RefundCheckDTO
}

// 初始化TaobaoRdcAligeniusRefundsCheckRequest对象
func NewTaobaoRdcAligeniusRefundsCheckRequest() *TaobaoRdcAligeniusRefundsCheckRequest{
    return &TaobaoRdcAligeniusRefundsCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusRefundsCheckRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.refunds.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusRefundsCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *TaobaoRdcAligeniusRefundsCheckRequest) SetParam(_param *RefundCheckDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoRdcAligeniusRefundsCheckRequest) GetParam() *RefundCheckDTO {
    return r._param
}
