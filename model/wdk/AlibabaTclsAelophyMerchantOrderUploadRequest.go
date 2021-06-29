package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家订单数据上传 API请求
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传
*/
type AlibabaTclsAelophyMerchantOrderUploadRequest struct {
    model.Params
    // 商家订单信息
    orderInfo   *MerchantOrderInfo
}

// 初始化AlibabaTclsAelophyMerchantOrderUploadRequest对象
func NewAlibabaTclsAelophyMerchantOrderUploadRequest() *AlibabaTclsAelophyMerchantOrderUploadRequest{
    return &AlibabaTclsAelophyMerchantOrderUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.order.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderInfo Setter
// 商家订单信息
func (r *AlibabaTclsAelophyMerchantOrderUploadRequest) SetOrderInfo(orderInfo *MerchantOrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlibabaTclsAelophyMerchantOrderUploadRequest) GetOrderInfo() *MerchantOrderInfo {
    return r.orderInfo
}
