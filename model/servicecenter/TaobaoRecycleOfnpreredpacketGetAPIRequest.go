package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商查询前置补贴红包的最新数据 API请求
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据
*/
type TaobaoRecycleOfnpreredpacketGetAPIRequest struct {
    model.Params
    // 旧机单id
    _oldOrderId   int64
}

// 初始化TaobaoRecycleOfnpreredpacketGetAPIRequest对象
func NewTaobaoRecycleOfnpreredpacketGetRequest() *TaobaoRecycleOfnpreredpacketGetAPIRequest{
    return &TaobaoRecycleOfnpreredpacketGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetApiMethodName() string {
    return "taobao.recycle.ofnpreredpacket.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OldOrderId Setter
// 旧机单id
func (r *TaobaoRecycleOfnpreredpacketGetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
    r._oldOrderId = _oldOrderId
    r.Set("old_order_id", _oldOrderId)
    return nil
}

// OldOrderId Getter
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetOldOrderId() int64 {
    return r._oldOrderId
}
