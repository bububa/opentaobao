package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcaseupdateproductAPIRequest 替换橱窗商品 API请求
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
type AlibabascbpshowcaseupdateproductAPIRequest struct {
	model.Params
	// 橱窗id
	_windowId int64
	// 新的商品id
	_newProductId int64
}

// NewAlibabascbpshowcaseupdateproductRequest 初始化AlibabascbpshowcaseupdateproductAPIRequest对象
func NewAlibabascbpshowcaseupdateproductRequest() *AlibabascbpshowcaseupdateproductAPIRequest {
	return &AlibabascbpshowcaseupdateproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcaseupdateproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.updateproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcaseupdateproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcaseupdateproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWindowId is WindowId Setter
// 橱窗id
func (r *AlibabascbpshowcaseupdateproductAPIRequest) SetWindowId(_windowId int64) error {
	r._windowId = _windowId
	r.Set("window_id", _windowId)
	return nil
}

// GetWindowId WindowId Getter
func (r AlibabascbpshowcaseupdateproductAPIRequest) GetWindowId() int64 {
	return r._windowId
}

// SetNewProductId is NewProductId Setter
// 新的商品id
func (r *AlibabascbpshowcaseupdateproductAPIRequest) SetNewProductId(_newProductId int64) error {
	r._newProductId = _newProductId
	r.Set("new_product_id", _newProductId)
	return nil
}

// GetNewProductId NewProductId Getter
func (r AlibabascbpshowcaseupdateproductAPIRequest) GetNewProductId() int64 {
	return r._newProductId
}
