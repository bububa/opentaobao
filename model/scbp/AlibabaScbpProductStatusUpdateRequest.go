package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改P4P产品推广状态 API请求
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
type AlibabaScbpProductStatusUpdateRequest struct {
    model.Params
    // 产品ID列表
    _productIdList   []int64
    // enabled:开启,disabled:暂停
    _status   string
}

// 初始化AlibabaScbpProductStatusUpdateRequest对象
func NewAlibabaScbpProductStatusUpdateRequest() *AlibabaScbpProductStatusUpdateRequest{
    return &AlibabaScbpProductStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductIdList Setter
// 产品ID列表
func (r *AlibabaScbpProductStatusUpdateRequest) SetProductIdList(_productIdList []int64) error {
    r._productIdList = _productIdList
    r.Set("product_id_list", _productIdList)
    return nil
}

// ProductIdList Getter
func (r AlibabaScbpProductStatusUpdateRequest) GetProductIdList() []int64 {
    return r._productIdList
}
// Status Setter
// enabled:开启,disabled:暂停
func (r *AlibabaScbpProductStatusUpdateRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaScbpProductStatusUpdateRequest) GetStatus() string {
    return r._status
}
