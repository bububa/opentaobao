package icbuseller

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站服务市场订单详情接口 APIRequest
alibaba.seller.vendor.order.detail

国际站服务市场订单列表接口
*/
type AlibabaSellerVendorOrderDetailRequest struct {
    model.Params

    // 订单编号
    orderNo   string 

}

func NewAlibabaSellerVendorOrderDetailRequest() *AlibabaSellerVendorOrderDetailRequest{
    return &AlibabaSellerVendorOrderDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSellerVendorOrderDetailRequest) GetApiMethodName() string {
    return "alibaba.seller.vendor.order.detail"
}

func (r AlibabaSellerVendorOrderDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSellerVendorOrderDetailRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r AlibabaSellerVendorOrderDetailRequest) GetOrderNo() string {
    return r.orderNo
}

