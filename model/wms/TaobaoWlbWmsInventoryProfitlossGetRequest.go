package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单列表批量获取库存损益单信息 APIRequest
taobao.wlb.wms.inventory.profitloss.get

通过订单列表批量获取库存损益单信息
*/
type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    model.Params

    // 菜鸟订单编码
    cnOrderCode   string 

}

func NewTaobaoWlbWmsInventoryProfitlossGetRequest() *TaobaoWlbWmsInventoryProfitlossGetRequest{
    return &TaobaoWlbWmsInventoryProfitlossGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.profitloss.get"
}

func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsInventoryProfitlossGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

func (r TaobaoWlbWmsInventoryProfitlossGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}

