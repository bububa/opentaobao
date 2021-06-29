package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售货机订单物流信息回传 API请求
alibaba.lst.vending.order.update

零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
*/
type AlibabaLstVendingOrderUpdateRequest struct {
    model.Params
    // 零售通设备订单
    _vendingOrderDTO   *VendingOrderDTO
}

// 初始化AlibabaLstVendingOrderUpdateRequest对象
func NewAlibabaLstVendingOrderUpdateRequest() *AlibabaLstVendingOrderUpdateRequest{
    return &AlibabaLstVendingOrderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingOrderUpdateRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.order.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingOrderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VendingOrderDTO Setter
// 零售通设备订单
func (r *AlibabaLstVendingOrderUpdateRequest) SetVendingOrderDTO(_vendingOrderDTO *VendingOrderDTO) error {
    r._vendingOrderDTO = _vendingOrderDTO
    r.Set("vending_order_d_t_o", _vendingOrderDTO)
    return nil
}

// VendingOrderDTO Getter
func (r AlibabaLstVendingOrderUpdateRequest) GetVendingOrderDTO() *VendingOrderDTO {
    return r._vendingOrderDTO
}
