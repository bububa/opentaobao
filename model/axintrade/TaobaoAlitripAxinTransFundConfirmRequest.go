package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认资金单 API请求
taobao.alitrip.axin.trans.fund.confirm

通过外部订单号进行资金结算
*/
type TaobaoAlitripAxinTransFundConfirmRequest struct {
    model.Params
    // 外部订单编号
    _outerOrderId   string
}

// 初始化TaobaoAlitripAxinTransFundConfirmRequest对象
func NewTaobaoAlitripAxinTransFundConfirmRequest() *TaobaoAlitripAxinTransFundConfirmRequest{
    return &TaobaoAlitripAxinTransFundConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterOrderId Setter
// 外部订单编号
func (r *TaobaoAlitripAxinTransFundConfirmRequest) SetOuterOrderId(_outerOrderId string) error {
    r._outerOrderId = _outerOrderId
    r.Set("outer_order_id", _outerOrderId)
    return nil
}

// OuterOrderId Getter
func (r TaobaoAlitripAxinTransFundConfirmRequest) GetOuterOrderId() string {
    return r._outerOrderId
}
