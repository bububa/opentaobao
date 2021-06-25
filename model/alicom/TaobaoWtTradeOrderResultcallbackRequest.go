package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调接口 APIRequest
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
*/
type TaobaoWtTradeOrderResultcallbackRequest struct {
    model.Params

    // 系统自动生成
    param0   *OrderResultDto 

}

func NewTaobaoWtTradeOrderResultcallbackRequest() *TaobaoWtTradeOrderResultcallbackRequest{
    return &TaobaoWtTradeOrderResultcallbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWtTradeOrderResultcallbackRequest) GetApiMethodName() string {
    return "taobao.wt.trade.order.resultcallback"
}

func (r TaobaoWtTradeOrderResultcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWtTradeOrderResultcallbackRequest) SetParam0(param0 *OrderResultDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoWtTradeOrderResultcallbackRequest) GetParam0() *OrderResultDto {
    return r.param0
}

