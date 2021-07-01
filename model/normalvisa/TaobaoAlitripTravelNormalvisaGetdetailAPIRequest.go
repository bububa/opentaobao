package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔订单的详情 API请求
taobao.alitrip.travel.normalvisa.getdetail

获取单笔签证的详细记录
*/
type TaobaoAlitripTravelNormalvisaGetdetailAPIRequest struct {
    model.Params
    // 订单id
    _bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaGetdetailAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaGetdetailRequest() *TaobaoAlitripTravelNormalvisaGetdetailAPIRequest{
    return &TaobaoAlitripTravelNormalvisaGetdetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetdetailAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetdetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaGetdetailAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaGetdetailAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
