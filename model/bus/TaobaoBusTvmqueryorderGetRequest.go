package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机查询订单信息 APIRequest
taobao.bus.tvmqueryorder.get

查询订单详情
*/
type TaobaoBusTvmqueryorderGetRequest struct {
    model.Params

    // 阿里订单标编号
    alitripOrderId   string 

}

func NewTaobaoBusTvmqueryorderGetRequest() *TaobaoBusTvmqueryorderGetRequest{
    return &TaobaoBusTvmqueryorderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmqueryorderGetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmqueryorder.get"
}

func (r TaobaoBusTvmqueryorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmqueryorderGetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusTvmqueryorderGetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

