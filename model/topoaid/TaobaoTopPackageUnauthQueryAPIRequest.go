package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageUnauthQueryAPIRequest 查询某手机号下的包裹 API请求
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
type TaobaoTopPackageUnauthQueryAPIRequest struct {
	model.Params
	// 查询包裹请求
	_queryPackageListRequest *QueryPackageListRequest
}

// NewTaobaoTopPackageUnauthQueryRequest 初始化TaobaoTopPackageUnauthQueryAPIRequest对象
func NewTaobaoTopPackageUnauthQueryRequest() *TaobaoTopPackageUnauthQueryAPIRequest {
	return &TaobaoTopPackageUnauthQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopPackageUnauthQueryAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.unauth.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopPackageUnauthQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopPackageUnauthQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryPackageListRequest is QueryPackageListRequest Setter
// 查询包裹请求
func (r *TaobaoTopPackageUnauthQueryAPIRequest) SetQueryPackageListRequest(_queryPackageListRequest *QueryPackageListRequest) error {
	r._queryPackageListRequest = _queryPackageListRequest
	r.Set("query_package_list_request", _queryPackageListRequest)
	return nil
}

// GetQueryPackageListRequest QueryPackageListRequest Getter
func (r TaobaoTopPackageUnauthQueryAPIRequest) GetQueryPackageListRequest() *QueryPackageListRequest {
	return r._queryPackageListRequest
}
