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
type TaobaoBusRefundfeeGetRequest struct {
    model.Params
    // 飞猪订单号
    _aliTripOrderId   string
    // 飞猪子订单号
    _subOrderIds   []int64
}

// 初始化TaobaoBusRefundfeeGetRequest对象
func NewTaobaoBusRefundfeeGetRequest() *TaobaoBusRefundfeeGetRequest{
    return &TaobaoBusRefundfeeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundfeeGetRequest) GetApiMethodName() string {
    return "taobao.bus.refundfee.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundfeeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliTripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusRefundfeeGetRequest) SetAliTripOrderId(_aliTripOrderId string) error {
    r._aliTripOrderId = _aliTripOrderId
    r.Set("ali_trip_order_id", _aliTripOrderId)
    return nil
}

// AliTripOrderId Getter
func (r TaobaoBusRefundfeeGetRequest) GetAliTripOrderId() string {
    return r._aliTripOrderId
}
// SubOrderIds Setter
// 飞猪子订单号
func (r *TaobaoBusRefundfeeGetRequest) SetSubOrderIds(_subOrderIds []int64) error {
    r._subOrderIds = _subOrderIds
    r.Set("sub_order_ids", _subOrderIds)
    return nil
}

// SubOrderIds Getter
func (r TaobaoBusRefundfeeGetRequest) GetSubOrderIds() []int64 {
    return r._subOrderIds
}
