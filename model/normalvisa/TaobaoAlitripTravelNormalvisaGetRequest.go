package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取签证记录 API请求
taobao.alitrip.travel.normalvisa.get

用于获取普通签证的记录信息
*/
type TaobaoAlitripTravelNormalvisaGetRequest struct {
    model.Params
    // 订单号
    bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaGetRequest对象
func NewTaobaoAlitripTravelNormalvisaGetRequest() *TaobaoAlitripTravelNormalvisaGetRequest{
    return &TaobaoAlitripTravelNormalvisaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
