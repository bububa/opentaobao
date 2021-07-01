package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商客户关联信息 API请求
alibaba.seller.vendor.service.vendorprocess

服务商客户关联信息
*/
type AlibabaSellerVendorServiceVendorprocessAPIRequest struct {
    model.Params
    // order_num
    _orderNum   string
}

// 初始化AlibabaSellerVendorServiceVendorprocessAPIRequest对象
func NewAlibabaSellerVendorServiceVendorprocessRequest() *AlibabaSellerVendorServiceVendorprocessAPIRequest{
    return &AlibabaSellerVendorServiceVendorprocessAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.service.vendorprocess"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNum Setter
// order_num
func (r *AlibabaSellerVendorServiceVendorprocessAPIRequest) SetOrderNum(_orderNum string) error {
    r._orderNum = _orderNum
    r.Set("order_num", _orderNum)
    return nil
}

// OrderNum Getter
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetOrderNum() string {
    return r._orderNum
}
