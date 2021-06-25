package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消订单 APIRequest
taobao.bus.cancleorder.set

取消订单
*/
type TaobaoBusCancleorderSetRequest struct {
    model.Params

    // 阿里订单号
    aliOrderId   string 

}

func NewTaobaoBusCancleorderSetRequest() *TaobaoBusCancleorderSetRequest{
    return &TaobaoBusCancleorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusCancleorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.cancleorder.set"
}

func (r TaobaoBusCancleorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusCancleorderSetRequest) SetAliOrderId(aliOrderId string) error {
    r.aliOrderId = aliOrderId
    r.Set("ali_order_id", aliOrderId)
    return nil
}

func (r TaobaoBusCancleorderSetRequest) GetAliOrderId() string {
    return r.aliOrderId
}

