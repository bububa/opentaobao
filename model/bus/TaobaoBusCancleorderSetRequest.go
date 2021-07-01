package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消订单 API请求
taobao.bus.cancleorder.set

取消订单
*/
type TaobaoBusCancleorderSetAPIRequest struct {
    model.Params
    // 阿里订单号
    _aliOrderId   string
}

// 初始化TaobaoBusCancleorderSetAPIRequest对象
func NewTaobaoBusCancleorderSetRequest() *TaobaoBusCancleorderSetAPIRequest{
    return &TaobaoBusCancleorderSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusCancleorderSetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.cancleorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusCancleorderSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliOrderId Setter
// 阿里订单号
func (r *TaobaoBusCancleorderSetAPIRequest) SetAliOrderId(_aliOrderId string) error {
    r._aliOrderId = _aliOrderId
    r.Set("ali_order_id", _aliOrderId)
    return nil
}

// AliOrderId Getter
func (r TaobaoBusCancleorderSetAPIRequest) GetAliOrderId() string {
    return r._aliOrderId
}
