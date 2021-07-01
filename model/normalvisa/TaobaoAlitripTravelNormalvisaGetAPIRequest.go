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
type TaobaoAlitripTravelNormalvisaGetAPIRequest struct {
    model.Params
    // 订单号
    _bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaGetAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaGetRequest() *TaobaoAlitripTravelNormalvisaGetAPIRequest{
    return &TaobaoAlitripTravelNormalvisaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
