package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机查询订单信息 API请求
taobao.bus.tvmqueryorder.get

查询订单详情
*/
type TaobaoBusTvmqueryorderGetAPIRequest struct {
    model.Params
    // 阿里订单标编号
    _alitripOrderId   string
}

// 初始化TaobaoBusTvmqueryorderGetAPIRequest对象
func NewTaobaoBusTvmqueryorderGetRequest() *TaobaoBusTvmqueryorderGetAPIRequest{
    return &TaobaoBusTvmqueryorderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.tvmqueryorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 阿里订单标编号
func (r *TaobaoBusTvmqueryorderGetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmqueryorderGetAPIRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
