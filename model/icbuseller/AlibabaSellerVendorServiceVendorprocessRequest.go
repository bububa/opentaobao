package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商客户关联信息 APIRequest
alibaba.seller.vendor.service.vendorprocess

服务商客户关联信息
*/
type AlibabaSellerVendorServiceVendorprocessRequest struct {
    model.Params

    // order_num
    orderNum   string 

}

func NewAlibabaSellerVendorServiceVendorprocessRequest() *AlibabaSellerVendorServiceVendorprocessRequest{
    return &AlibabaSellerVendorServiceVendorprocessRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSellerVendorServiceVendorprocessRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.service.vendorprocess"
}

func (r AlibabaSellerVendorServiceVendorprocessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSellerVendorServiceVendorprocessRequest) SetOrderNum(orderNum string) error {
    r.orderNum = orderNum
    r.Set("order_num", orderNum)
    return nil
}

func (r AlibabaSellerVendorServiceVendorprocessRequest) GetOrderNum() string {
    return r.orderNum
}

