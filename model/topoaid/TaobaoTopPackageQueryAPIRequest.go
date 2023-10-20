package topoaid

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageQueryAPIRequest 淘系包裹查询 API请求
// taobao.top.package.query
//
// 淘系包裹查询
type TaobaoTopPackageQueryAPIRequest struct {
	model.Params
	// 查询用户的包裹列表
	_queryPackageListRequest *QueryPackageListRequest
}

// NewTaobaoTopPackageQueryRequest 初始化TaobaoTopPackageQueryAPIRequest对象
func NewTaobaoTopPackageQueryRequest() *TaobaoTopPackageQueryAPIRequest {
	return &TaobaoTopPackageQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopPackageQueryAPIRequest) Reset() {
	r._queryPackageListRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopPackageQueryAPIRequest) GetApiMethodName() string {
	return "taobao.top.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopPackageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopPackageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryPackageListRequest is QueryPackageListRequest Setter
// 查询用户的包裹列表
func (r *TaobaoTopPackageQueryAPIRequest) SetQueryPackageListRequest(_queryPackageListRequest *QueryPackageListRequest) error {
	r._queryPackageListRequest = _queryPackageListRequest
	r.Set("query_package_list_request", _queryPackageListRequest)
	return nil
}

// GetQueryPackageListRequest QueryPackageListRequest Getter
func (r TaobaoTopPackageQueryAPIRequest) GetQueryPackageListRequest() *QueryPackageListRequest {
	return r._queryPackageListRequest
}

var poolTaobaoTopPackageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopPackageQueryRequest()
	},
}

// GetTaobaoTopPackageQueryRequest 从 sync.Pool 获取 TaobaoTopPackageQueryAPIRequest
func GetTaobaoTopPackageQueryAPIRequest() *TaobaoTopPackageQueryAPIRequest {
	return poolTaobaoTopPackageQueryAPIRequest.Get().(*TaobaoTopPackageQueryAPIRequest)
}

// ReleaseTaobaoTopPackageQueryAPIRequest 将 TaobaoTopPackageQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopPackageQueryAPIRequest(v *TaobaoTopPackageQueryAPIRequest) {
	v.Reset()
	poolTaobaoTopPackageQueryAPIRequest.Put(v)
}
