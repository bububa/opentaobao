package cityretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城零售完美履约转仓 APIRequest
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓
*/
type TaobaoCityretailWmflConvertWarehouseRequest struct {
    model.Params

    // 淘宝交易单id
    tbOrderId   string 

}

func NewTaobaoCityretailWmflConvertWarehouseRequest() *TaobaoCityretailWmflConvertWarehouseRequest{
    return &TaobaoCityretailWmflConvertWarehouseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCityretailWmflConvertWarehouseRequest) GetApiMethodName() string {
    return "taobao.cityretail.wmfl.convert.warehouse"
}

func (r TaobaoCityretailWmflConvertWarehouseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCityretailWmflConvertWarehouseRequest) SetTbOrderId(tbOrderId string) error {
    r.tbOrderId = tbOrderId
    r.Set("tb_order_id", tbOrderId)
    return nil
}

func (r TaobaoCityretailWmflConvertWarehouseRequest) GetTbOrderId() string {
    return r.tbOrderId
}

