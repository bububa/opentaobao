package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退票费用明细 API请求
taobao.bus.refundfee.get

查询退票的费用信息
*/
type TaobaoBusRefundfeeGetAPIRequest struct {
    model.Params
    // 飞猪订单号
    _aliTripOrderId   string
    // 飞猪子订单号
    _subOrderIds   []int64
}

// 初始化TaobaoBusRefundfeeGetAPIRequest对象
func NewTaobaoBusRefundfeeGetRequest() *TaobaoBusRefundfeeGetAPIRequest{
    return &TaobaoBusRefundfeeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundfeeGetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.refundfee.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundfeeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliTripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusRefundfeeGetAPIRequest) SetAliTripOrderId(_aliTripOrderId string) error {
    r._aliTripOrderId = _aliTripOrderId
    r.Set("ali_trip_order_id", _aliTripOrderId)
    return nil
}

// AliTripOrderId Getter
func (r TaobaoBusRefundfeeGetAPIRequest) GetAliTripOrderId() string {
    return r._aliTripOrderId
}
// SubOrderIds Setter
// 飞猪子订单号
func (r *TaobaoBusRefundfeeGetAPIRequest) SetSubOrderIds(_subOrderIds []int64) error {
    r._subOrderIds = _subOrderIds
    r.Set("sub_order_ids", _subOrderIds)
    return nil
}

// SubOrderIds Getter
func (r TaobaoBusRefundfeeGetAPIRequest) GetSubOrderIds() []int64 {
    return r._subOrderIds
}
