package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商客户关联信息 APIRequest
alibaba.seller.vendor.service.process

服务商客户关联信息
*/
type AlibabaSellerVendorServiceProcessRequest struct {
    model.Params

    // order_num
    orderNum   string 

}

func NewAlibabaSellerVendorServiceProcessRequest() *AlibabaSellerVendorServiceProcessRequest{
    return &AlibabaSellerVendorServiceProcessRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSellerVendorServiceProcessRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.service.process"
}

func (r AlibabaSellerVendorServiceProcessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSellerVendorServiceProcessRequest) SetOrderNum(orderNum string) error {
    r.orderNum = orderNum
    r.Set("order_num", orderNum)
    return nil
}

func (r AlibabaSellerVendorServiceProcessRequest) GetOrderNum() string {
    return r.orderNum
}

