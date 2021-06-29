package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-处罚订单查询 API请求
taobao.tbk.dg.punish.order.get

新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
*/
type TaobaoTbkDgPunishOrderGetRequest struct {
    model.Params
    // 入参的对象
    _afOrderOption   *TopApiAfOrderOption
}

// 初始化TaobaoTbkDgPunishOrderGetRequest对象
func NewTaobaoTbkDgPunishOrderGetRequest() *TaobaoTbkDgPunishOrderGetRequest{
    return &TaobaoTbkDgPunishOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgPunishOrderGetRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.punish.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgPunishOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AfOrderOption Setter
// 入参的对象
func (r *TaobaoTbkDgPunishOrderGetRequest) SetAfOrderOption(_afOrderOption *TopApiAfOrderOption) error {
    r._afOrderOption = _afOrderOption
    r.Set("af_order_option", _afOrderOption)
    return nil
}

// AfOrderOption Getter
func (r TaobaoTbkDgPunishOrderGetRequest) GetAfOrderOption() *TopApiAfOrderOption {
    return r._afOrderOption
}
