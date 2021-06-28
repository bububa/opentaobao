package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改P4P产品推广状态 APIRequest
alibaba.scbp.product.status.update

修改P4P产品推广状态
*/
type AlibabaScbpProductStatusUpdateRequest struct {
    model.Params

    // 产品ID列表
    productIdList   []int64 

    // enabled:开启,disabled:暂停
    status   string 

}

func NewAlibabaScbpProductStatusUpdateRequest() *AlibabaScbpProductStatusUpdateRequest{
    return &AlibabaScbpProductStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpProductStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.status.update"
}

func (r AlibabaScbpProductStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpProductStatusUpdateRequest) SetProductIdList(productIdList []int64) error {
    r.productIdList = productIdList
    r.Set("product_id_list", productIdList)
    return nil
}

func (r AlibabaScbpProductStatusUpdateRequest) GetProductIdList() []int64 {
    return r.productIdList
}

func (r *AlibabaScbpProductStatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaScbpProductStatusUpdateRequest) GetStatus() string {
    return r.status
}

