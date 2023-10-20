package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcasedeleteproductAPIRequest 批量删除橱窗商品 API请求
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
type AlibabascbpshowcasedeleteproductAPIRequest struct {
	model.Params
	// 橱窗idList
	_windowIdList []string
}

// NewAlibabascbpshowcasedeleteproductRequest 初始化AlibabascbpshowcasedeleteproductAPIRequest对象
func NewAlibabascbpshowcasedeleteproductRequest() *AlibabascbpshowcasedeleteproductAPIRequest {
	return &AlibabascbpshowcasedeleteproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcasedeleteproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.deleteproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcasedeleteproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcasedeleteproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWindowIdList is WindowIdList Setter
// 橱窗idList
func (r *AlibabascbpshowcasedeleteproductAPIRequest) SetWindowIdList(_windowIdList []string) error {
	r._windowIdList = _windowIdList
	r.Set("window_id_list", _windowIdList)
	return nil
}

// GetWindowIdList WindowIdList Getter
func (r AlibabascbpshowcasedeleteproductAPIRequest) GetWindowIdList() []string {
	return r._windowIdList
}
