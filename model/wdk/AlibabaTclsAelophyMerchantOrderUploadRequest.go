package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家订单数据上传 APIRequest
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传
*/
type AlibabaTclsAelophyMerchantOrderUploadRequest struct {
    model.Params

    // 商家订单信息
    orderInfo   *MerchantOrderInfo 

}

func NewAlibabaTclsAelophyMerchantOrderUploadRequest() *AlibabaTclsAelophyMerchantOrderUploadRequest{
    return &AlibabaTclsAelophyMerchantOrderUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.order.upload"
}

func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantOrderUploadRequest) SetOrderInfo(orderInfo *MerchantOrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetOrderInfo() *MerchantOrderInfo {
    return r.orderInfo
}

