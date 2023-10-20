package icbushowcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseUpdateproductAPIRequest 替换橱窗商品 API请求
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
type AlibabaScbpShowcaseUpdateproductAPIRequest struct {
	model.Params
	// 橱窗id
	_windowId int64
	// 新的商品id
	_newProductId int64
}

// NewAlibabaScbpShowcaseUpdateproductRequest 初始化AlibabaScbpShowcaseUpdateproductAPIRequest对象
func NewAlibabaScbpShowcaseUpdateproductRequest() *AlibabaScbpShowcaseUpdateproductAPIRequest {
	return &AlibabaScbpShowcaseUpdateproductAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) Reset() {
	r._windowId = 0
	r._newProductId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.updateproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWindowId is WindowId Setter
// 橱窗id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetWindowId(_windowId int64) error {
	r._windowId = _windowId
	r.Set("window_id", _windowId)
	return nil
}

// GetWindowId WindowId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetWindowId() int64 {
	return r._windowId
}

// SetNewProductId is NewProductId Setter
// 新的商品id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetNewProductId(_newProductId int64) error {
	r._newProductId = _newProductId
	r.Set("new_product_id", _newProductId)
	return nil
}

// GetNewProductId NewProductId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetNewProductId() int64 {
	return r._newProductId
}

var poolAlibabaScbpShowcaseUpdateproductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseUpdateproductRequest()
	},
}

// GetAlibabaScbpShowcaseUpdateproductRequest 从 sync.Pool 获取 AlibabaScbpShowcaseUpdateproductAPIRequest
func GetAlibabaScbpShowcaseUpdateproductAPIRequest() *AlibabaScbpShowcaseUpdateproductAPIRequest {
	return poolAlibabaScbpShowcaseUpdateproductAPIRequest.Get().(*AlibabaScbpShowcaseUpdateproductAPIRequest)
}

// ReleaseAlibabaScbpShowcaseUpdateproductAPIRequest 将 AlibabaScbpShowcaseUpdateproductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseUpdateproductAPIRequest(v *AlibabaScbpShowcaseUpdateproductAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseUpdateproductAPIRequest.Put(v)
}
