package cityretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城零售完美履约转仓 API请求
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓
*/
type TaobaoCityretailWmflConvertWarehouseRequest struct {
    model.Params
    // 淘宝交易单id
    _tbOrderId   string
}

// 初始化TaobaoCityretailWmflConvertWarehouseRequest对象
func NewTaobaoCityretailWmflConvertWarehouseRequest() *TaobaoCityretailWmflConvertWarehouseRequest{
    return &TaobaoCityretailWmflConvertWarehouseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCityretailWmflConvertWarehouseRequest) GetApiMethodName() string {
    return "taobao.cityretail.wmfl.convert.warehouse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCityretailWmflConvertWarehouseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbOrderId Setter
// 淘宝交易单id
func (r *TaobaoCityretailWmflConvertWarehouseRequest) SetTbOrderId(_tbOrderId string) error {
    r._tbOrderId = _tbOrderId
    r.Set("tb_order_id", _tbOrderId)
    return nil
}

// TbOrderId Getter
func (r TaobaoCityretailWmflConvertWarehouseRequest) GetTbOrderId() string {
    return r._tbOrderId
}
