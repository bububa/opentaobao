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
type TaobaoAlitripTravelNormalvisaGetdetailRequest struct {
    model.Params
    // 订单id
    bizOrderId   int64
}

// 初始化TaobaoAlitripTravelNormalvisaGetdetailRequest对象
func NewTaobaoAlitripTravelNormalvisaGetdetailRequest() *TaobaoAlitripTravelNormalvisaGetdetailRequest{
    return &TaobaoAlitripTravelNormalvisaGetdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaGetdetailRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaGetdetailRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
