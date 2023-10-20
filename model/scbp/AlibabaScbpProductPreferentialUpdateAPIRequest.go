package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductPreferentialUpdateAPIRequest 设置P4P产品优先推广状态 API请求
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
type AlibabaScbpProductPreferentialUpdateAPIRequest struct {
	model.Params
	// Y:设置优推,N:取消优推
	_status string
	// 关键词ID
	_keywordId int64
	// 产品ID
	_productId int64
}

// NewAlibabaScbpProductPreferentialUpdateRequest 初始化AlibabaScbpProductPreferentialUpdateAPIRequest对象
func NewAlibabaScbpProductPreferentialUpdateRequest() *AlibabaScbpProductPreferentialUpdateAPIRequest {
	return &AlibabaScbpProductPreferentialUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) Reset() {
	r._status = ""
	r._keywordId = 0
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.preferential.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// Y:设置优推,N:取消优推
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetStatus() string {
	return r._status
}

// SetKeywordId is KeywordId Setter
// 关键词ID
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetKeywordId(_keywordId int64) error {
	r._keywordId = _keywordId
	r.Set("keyword_id", _keywordId)
	return nil
}

// GetKeywordId KeywordId Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetKeywordId() int64 {
	return r._keywordId
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAlibabaScbpProductPreferentialUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpProductPreferentialUpdateRequest()
	},
}

// GetAlibabaScbpProductPreferentialUpdateRequest 从 sync.Pool 获取 AlibabaScbpProductPreferentialUpdateAPIRequest
func GetAlibabaScbpProductPreferentialUpdateAPIRequest() *AlibabaScbpProductPreferentialUpdateAPIRequest {
	return poolAlibabaScbpProductPreferentialUpdateAPIRequest.Get().(*AlibabaScbpProductPreferentialUpdateAPIRequest)
}

// ReleaseAlibabaScbpProductPreferentialUpdateAPIRequest 将 AlibabaScbpProductPreferentialUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpProductPreferentialUpdateAPIRequest(v *AlibabaScbpProductPreferentialUpdateAPIRequest) {
	v.Reset()
	poolAlibabaScbpProductPreferentialUpdateAPIRequest.Put(v)
}
