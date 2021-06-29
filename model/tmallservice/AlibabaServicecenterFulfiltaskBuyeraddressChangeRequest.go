package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改消费者服务地址 API请求
alibaba.servicecenter.fulfiltask.buyeraddress.change

当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
*/
type AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest struct {
    model.Params
    // 核销单id
    fulfilTaskId   int64
    // 详细地址
    addressDetail   string
    // 地址编码
    location   int64
}

// 初始化AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest对象
func NewAlibabaServicecenterFulfiltaskBuyeraddressChangeRequest() *AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest{
    return &AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.fulfiltask.buyeraddress.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfilTaskId Setter
// 核销单id
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}
// AddressDetail Setter
// 详细地址
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) SetAddressDetail(addressDetail string) error {
    r.addressDetail = addressDetail
    r.Set("address_detail", addressDetail)
    return nil
}

// AddressDetail Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) GetAddressDetail() string {
    return r.addressDetail
}
// Location Setter
// 地址编码
func (r *AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) SetLocation(location int64) error {
    r.location = location
    r.Set("location", location)
    return nil
}

// Location Getter
func (r AlibabaServicecenterFulfiltaskBuyeraddressChangeRequest) GetLocation() int64 {
    return r.location
}
