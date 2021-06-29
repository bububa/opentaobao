package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取签证记录 APIRequest
taobao.alitrip.travel.normalvisa.get

用于获取普通签证的记录信息
*/
type TaobaoAlitripTravelNormalvisaGetRequest struct {
    model.Params

    // 订单号
    bizOrderId   int64 

}

func NewTaobaoAlitripTravelNormalvisaGetRequest() *TaobaoAlitripTravelNormalvisaGetRequest{
    return &TaobaoAlitripTravelNormalvisaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.get"
}

func (r TaobaoAlitripTravelNormalvisaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

