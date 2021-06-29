package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
已入塔商家查询订单列表 APIRequest
tmall.nr.sold.orderlist.query.jst

该服务用于已入聚石塔的商家，获取订单列表信息；
*/
type TmallNrSoldOrderlistQueryJstRequest struct {
    model.Params

    // 入参对象
    param0   *NrTimingOrderSoldQueryReqDto 

}

func NewTmallNrSoldOrderlistQueryJstRequest() *TmallNrSoldOrderlistQueryJstRequest{
    return &TmallNrSoldOrderlistQueryJstRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrSoldOrderlistQueryJstRequest) GetApiMethodName() string {
    return "tmall.nr.sold.orderlist.query.jst"
}

func (r TmallNrSoldOrderlistQueryJstRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrSoldOrderlistQueryJstRequest) SetParam0(param0 *NrTimingOrderSoldQueryReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrSoldOrderlistQueryJstRequest) GetParam0() *NrTimingOrderSoldQueryReqDto {
    return r.param0
}

