package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGetAPIRequest 获得单个商品详情 API请求
// alibaba.icbu.product.get
//
// 获取商品详情
type AlibabaIcbuProductGetAPIRequest struct {
	model.Params
	// 商品语种，目前只支持ENGLISH
	_language string
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaIcbuProductGetRequest 初始化AlibabaIcbuProductGetAPIRequest对象
func NewAlibabaIcbuProductGetRequest() *AlibabaIcbuProductGetAPIRequest {
	return &AlibabaIcbuProductGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductGetAPIRequest) Reset() {
	r._language = ""
	r._productId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 商品语种，目前只支持ENGLISH
func (r *AlibabaIcbuProductGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductGetAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductGetAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaIcbuProductGetAPIRequest) GetProductId() string {
	return r._productId
}

var poolAlibabaIcbuProductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductGetRequest()
	},
}

// GetAlibabaIcbuProductGetRequest 从 sync.Pool 获取 AlibabaIcbuProductGetAPIRequest
func GetAlibabaIcbuProductGetAPIRequest() *AlibabaIcbuProductGetAPIRequest {
	return poolAlibabaIcbuProductGetAPIRequest.Get().(*AlibabaIcbuProductGetAPIRequest)
}

// ReleaseAlibabaIcbuProductGetAPIRequest 将 AlibabaIcbuProductGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductGetAPIRequest(v *AlibabaIcbuProductGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductGetAPIRequest.Put(v)
}
