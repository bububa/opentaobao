package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenPresalespackageConsignAPIRequest 预售预包尾款推单发货 API请求
// taobao.qimen.presalespackage.consign
//
// 预售预包尾款推单发货
type TaobaoQimenPresalespackageConsignAPIRequest struct {
	model.Params
	// 请求
	_request *PresalesPackageConsignRequest
}

// NewTaobaoQimenPresalespackageConsignRequest 初始化TaobaoQimenPresalespackageConsignAPIRequest对象
func NewTaobaoQimenPresalespackageConsignRequest() *TaobaoQimenPresalespackageConsignAPIRequest {
	return &TaobaoQimenPresalespackageConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenPresalespackageConsignAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.presalespackage.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenPresalespackageConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenPresalespackageConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 请求
func (r *TaobaoQimenPresalespackageConsignAPIRequest) SetRequest(_request *PresalesPackageConsignRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenPresalespackageConsignAPIRequest) GetRequest() *PresalesPackageConsignRequest {
	return r._request
}
