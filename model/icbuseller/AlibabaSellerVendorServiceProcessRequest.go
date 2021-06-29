package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商客户关联信息 API请求
alibaba.seller.vendor.service.process

服务商客户关联信息
*/
type AlibabaSellerVendorServiceProcessRequest struct {
    model.Params
    // order_num
    orderNum   string
}

// 初始化AlibabaSellerVendorServiceProcessRequest对象
func NewAlibabaSellerVendorServiceProcessRequest() *AlibabaSellerVendorServiceProcessRequest{
    return &AlibabaSellerVendorServiceProcessRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceProcessRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.service.process"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceProcessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNum Setter
// order_num
func (r *AlibabaSellerVendorServiceProcessRequest) SetOrderNum(orderNum string) error {
    r.orderNum = orderNum
    r.Set("order_num", orderNum)
    return nil
}

// OrderNum Getter
func (r AlibabaSellerVendorServiceProcessRequest) GetOrderNum() string {
    return r.orderNum
}
