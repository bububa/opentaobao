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
type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    model.Params
    // 菜鸟订单编码
    cnOrderCode   string
}

// 初始化TaobaoWlbWmsInventoryProfitlossGetRequest对象
func NewTaobaoWlbWmsInventoryProfitlossGetRequest() *TaobaoWlbWmsInventoryProfitlossGetRequest{
    return &TaobaoWlbWmsInventoryProfitlossGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.profitloss.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CnOrderCode Setter
// 菜鸟订单编码
func (r *TaobaoWlbWmsInventoryProfitlossGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}
