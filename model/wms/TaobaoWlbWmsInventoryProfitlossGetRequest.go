package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单列表批量获取库存损益单信息 API请求
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息
*/
type TaobaoWlbWmsInventoryProfitlossGetAPIRequest struct {
    model.Params
    // 菜鸟订单编码
    _cnOrderCode   string
}

// 初始化TaobaoWlbWmsInventoryProfitlossGetAPIRequest对象
func NewTaobaoWlbWmsInventoryProfitlossGetRequest() *TaobaoWlbWmsInventoryProfitlossGetAPIRequest{
    return &TaobaoWlbWmsInventoryProfitlossGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.profitloss.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CnOrderCode Setter
// 菜鸟订单编码
func (r *TaobaoWlbWmsInventoryProfitlossGetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
    r._cnOrderCode = _cnOrderCode
    r.Set("cn_order_code", _cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetCnOrderCode() string {
    return r._cnOrderCode
}
