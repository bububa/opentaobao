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
    vendingOrderDTO   *VendingOrderDto
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
func (r *AlibabaLstVendingOrderUpdateRequest) SetVendingOrderDTO(vendingOrderDTO *VendingOrderDto) error {
    r.vendingOrderDTO = vendingOrderDTO
    r.Set("vending_order_d_t_o", vendingOrderDTO)
    return nil
}

// VendingOrderDTO Getter
func (r AlibabaLstVendingOrderUpdateRequest) GetVendingOrderDTO() *VendingOrderDto {
    return r.vendingOrderDTO
}
