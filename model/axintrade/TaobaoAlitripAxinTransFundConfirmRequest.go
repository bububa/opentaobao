package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认资金单 APIRequest
taobao.alitrip.axin.trans.fund.confirm

通过外部订单号进行资金结算
*/
type TaobaoAlitripAxinTransFundConfirmRequest struct {
    model.Params

    // 外部订单编号
    outerOrderId   string 

}

func NewTaobaoAlitripAxinTransFundConfirmRequest() *TaobaoAlitripAxinTransFundConfirmRequest{
    return &TaobaoAlitripAxinTransFundConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransFundConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.confirm"
}

func (r TaobaoAlitripAxinTransFundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransFundConfirmRequest) SetOuterOrderId(outerOrderId string) error {
    r.outerOrderId = outerOrderId
    r.Set("outer_order_id", outerOrderId)
    return nil
}

func (r TaobaoAlitripAxinTransFundConfirmRequest) GetOuterOrderId() string {
    return r.outerOrderId
}

