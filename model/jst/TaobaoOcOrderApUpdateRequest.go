package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按OC订单分账 API请求
taobao.oc.order.ap.update

对OC订单执行分账操作
*/
type TaobaoOcOrderApUpdateRequest struct {
    model.Params
    // 调用创建OC订单接口生成的id
    _ocOrderId   int64
}

// 初始化TaobaoOcOrderApUpdateRequest对象
func NewTaobaoOcOrderApUpdateRequest() *TaobaoOcOrderApUpdateRequest{
    return &TaobaoOcOrderApUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcOrderApUpdateRequest) GetApiMethodName() string {
    return "taobao.oc.order.ap.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcOrderApUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OcOrderId Setter
// 调用创建OC订单接口生成的id
func (r *TaobaoOcOrderApUpdateRequest) SetOcOrderId(_ocOrderId int64) error {
    r._ocOrderId = _ocOrderId
    r.Set("oc_order_id", _ocOrderId)
    return nil
}

// OcOrderId Getter
func (r TaobaoOcOrderApUpdateRequest) GetOcOrderId() int64 {
    return r._ocOrderId
}
