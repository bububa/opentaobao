package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商查询前置补贴红包的最新数据 APIRequest
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据
*/
type TaobaoRecycleOfnpreredpacketGetRequest struct {
    model.Params

    // 旧机单id
    oldOrderId   int64 

}

func NewTaobaoRecycleOfnpreredpacketGetRequest() *TaobaoRecycleOfnpreredpacketGetRequest{
    return &TaobaoRecycleOfnpreredpacketGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRecycleOfnpreredpacketGetRequest) GetApiMethodName() string {
    return "taobao.recycle.ofnpreredpacket.get"
}

func (r TaobaoRecycleOfnpreredpacketGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRecycleOfnpreredpacketGetRequest) SetOldOrderId(oldOrderId int64) error {
    r.oldOrderId = oldOrderId
    r.Set("old_order_id", oldOrderId)
    return nil
}

func (r TaobaoRecycleOfnpreredpacketGetRequest) GetOldOrderId() int64 {
    return r.oldOrderId
}

