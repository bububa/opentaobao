package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置P4P产品优先推广状态 APIRequest
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态
*/
type AlibabaScbpProductPreferentialUpdateRequest struct {
    model.Params

    // 关键词ID
    keywordId   int64 

    // 产品ID
    productId   int64 

    // Y:设置优推,N:取消优推
    status   string 

}

func NewAlibabaScbpProductPreferentialUpdateRequest() *AlibabaScbpProductPreferentialUpdateRequest{
    return &AlibabaScbpProductPreferentialUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpProductPreferentialUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.preferential.update"
}

func (r AlibabaScbpProductPreferentialUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpProductPreferentialUpdateRequest) SetKeywordId(keywordId int64) error {
    r.keywordId = keywordId
    r.Set("keyword_id", keywordId)
    return nil
}

func (r AlibabaScbpProductPreferentialUpdateRequest) GetKeywordId() int64 {
    return r.keywordId
}

func (r *AlibabaScbpProductPreferentialUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaScbpProductPreferentialUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *AlibabaScbpProductPreferentialUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaScbpProductPreferentialUpdateRequest) GetStatus() string {
    return r.status
}

