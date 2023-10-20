package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpproductpreferentialupdateAPIRequest 设置P4P产品优先推广状态 API请求
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
type AlibabascbpproductpreferentialupdateAPIRequest struct {
	model.Params
	// Y:设置优推,N:取消优推
	_status string
	// 关键词ID
	_keywordId int64
	// 产品ID
	_productId int64
}

// NewAlibabascbpproductpreferentialupdateRequest 初始化AlibabascbpproductpreferentialupdateAPIRequest对象
func NewAlibabascbpproductpreferentialupdateRequest() *AlibabascbpproductpreferentialupdateAPIRequest {
	return &AlibabascbpproductpreferentialupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.preferential.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// Y:设置优推,N:取消优推
func (r *AlibabascbpproductpreferentialupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetStatus() string {
	return r._status
}

// SetKeywordId is KeywordId Setter
// 关键词ID
func (r *AlibabascbpproductpreferentialupdateAPIRequest) SetKeywordId(_keywordId int64) error {
	r._keywordId = _keywordId
	r.Set("keyword_id", _keywordId)
	return nil
}

// GetKeywordId KeywordId Getter
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetKeywordId() int64 {
	return r._keywordId
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabascbpproductpreferentialupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabascbpproductpreferentialupdateAPIRequest) GetProductId() int64 {
	return r._productId
}
