package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置P4P产品优先推广状态 API请求
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

// 初始化AlibabaScbpProductPreferentialUpdateRequest对象
func NewAlibabaScbpProductPreferentialUpdateRequest() *AlibabaScbpProductPreferentialUpdateRequest{
    return &AlibabaScbpProductPreferentialUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductPreferentialUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.preferential.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductPreferentialUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordId Setter
// 关键词ID
func (r *AlibabaScbpProductPreferentialUpdateRequest) SetKeywordId(keywordId int64) error {
    r.keywordId = keywordId
    r.Set("keyword_id", keywordId)
    return nil
}

// KeywordId Getter
func (r AlibabaScbpProductPreferentialUpdateRequest) GetKeywordId() int64 {
    return r.keywordId
}
// ProductId Setter
// 产品ID
func (r *AlibabaScbpProductPreferentialUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaScbpProductPreferentialUpdateRequest) GetProductId() int64 {
    return r.productId
}
// Status Setter
// Y:设置优推,N:取消优推
func (r *AlibabaScbpProductPreferentialUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaScbpProductPreferentialUpdateRequest) GetStatus() string {
    return r.status
}
