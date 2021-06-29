package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔订单的详情 APIRequest
taobao.alitrip.travel.normalvisa.getdetail

获取单笔签证的详细记录
*/
type TaobaoAlitripTravelNormalvisaGetdetailRequest struct {
    model.Params

    // 订单id
    bizOrderId   int64 

}

func NewTaobaoAlitripTravelNormalvisaGetdetailRequest() *TaobaoAlitripTravelNormalvisaGetdetailRequest{
    return &TaobaoAlitripTravelNormalvisaGetdetailRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.getdetail"
}

func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaGetdetailRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

