package icbushowcase

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.updateproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WindowId Setter
// 橱窗id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetWindowId(_windowId int64) error {
	r._windowId = _windowId
	r.Set("window_id", _windowId)
	return nil
}

// Get WindowId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetWindowId() int64 {
	return r._windowId
}

// Set is NewProductId Setter
// 新的商品id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetNewProductId(_newProductId int64) error {
	r._newProductId = _newProductId
	r.Set("new_product_id", _newProductId)
	return nil
}

// Get NewProductId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetNewProductId() int64 {
	return r._newProductId
}
