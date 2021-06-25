package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-处罚订单查询 APIRequest
taobao.tbk.dg.punish.order.get

新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
*/
type TaobaoTbkDgPunishOrderGetRequest struct {
    model.Params

    // 入参的对象
    afOrderOption   *TopApiAfOrderOption 

}

func NewTaobaoTbkDgPunishOrderGetRequest() *TaobaoTbkDgPunishOrderGetRequest{
    return &TaobaoTbkDgPunishOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkDgPunishOrderGetRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.punish.order.get"
}

func (r TaobaoTbkDgPunishOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkDgPunishOrderGetRequest) SetAfOrderOption(afOrderOption *TopApiAfOrderOption) error {
    r.afOrderOption = afOrderOption
    r.Set("af_order_option", afOrderOption)
    return nil
}

func (r TaobaoTbkDgPunishOrderGetRequest) GetAfOrderOption() *TopApiAfOrderOption {
    return r.afOrderOption
}

