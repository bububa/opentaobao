package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest Dropshipper查找商品信息接口 API请求
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest struct {
	model.Params
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
	// 商品ID
	_productId int64
}

// NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest 初始化AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest对象
func NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest {
	return &AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) Reset() {
	r._localCountry = ""
	r._localLanguage = ""
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetApiMethodName() string {
	return "aliexpress.postproduct.redefining.findaeproductbyidfordropshipper"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocalCountry is LocalCountry Setter
// 国家
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetLocalCountry(_localCountry string) error {
	r._localCountry = _localCountry
	r.Set("local_country", _localCountry)
	return nil
}

// GetLocalCountry LocalCountry Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetLocalCountry() string {
	return r._localCountry
}

// SetLocalLanguage is LocalLanguage Setter
// 语言
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetLocalLanguage(_localLanguage string) error {
	r._localLanguage = _localLanguage
	r.Set("local_language", _localLanguage)
	return nil
}

// GetLocalLanguage LocalLanguage Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetLocalLanguage() string {
	return r._localLanguage
}

// SetProductId is ProductId Setter
// 商品ID
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest()
	},
}

// GetAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest 从 sync.Pool 获取 AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest
func GetAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest {
	return poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest.Get().(*AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest)
}

// ReleaseAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest 将 AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest 放入 sync.Pool
func ReleaseAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest(v *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) {
	v.Reset()
	poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest.Put(v)
}
