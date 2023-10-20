package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAppletPackageQueryAPIRequest 淘宝包裹查询 API请求
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
type TaobaoLogisticsAppletPackageQueryAPIRequest struct {
	model.Params
	// 查询用户的包裹列表
	_queryPackageListRequest *QueryPackageListRequest
}

// NewTaobaoLogisticsAppletPackageQueryRequest 初始化TaobaoLogisticsAppletPackageQueryAPIRequest对象
func NewTaobaoLogisticsAppletPackageQueryRequest() *TaobaoLogisticsAppletPackageQueryAPIRequest {
	return &TaobaoLogisticsAppletPackageQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsAppletPackageQueryAPIRequest) Reset() {
	r._queryPackageListRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAppletPackageQueryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.applet.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAppletPackageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsAppletPackageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryPackageListRequest is QueryPackageListRequest Setter
// 查询用户的包裹列表
func (r *TaobaoLogisticsAppletPackageQueryAPIRequest) SetQueryPackageListRequest(_queryPackageListRequest *QueryPackageListRequest) error {
	r._queryPackageListRequest = _queryPackageListRequest
	r.Set("query_package_list_request", _queryPackageListRequest)
	return nil
}

// GetQueryPackageListRequest QueryPackageListRequest Getter
func (r TaobaoLogisticsAppletPackageQueryAPIRequest) GetQueryPackageListRequest() *QueryPackageListRequest {
	return r._queryPackageListRequest
}

var poolTaobaoLogisticsAppletPackageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsAppletPackageQueryRequest()
	},
}

// GetTaobaoLogisticsAppletPackageQueryRequest 从 sync.Pool 获取 TaobaoLogisticsAppletPackageQueryAPIRequest
func GetTaobaoLogisticsAppletPackageQueryAPIRequest() *TaobaoLogisticsAppletPackageQueryAPIRequest {
	return poolTaobaoLogisticsAppletPackageQueryAPIRequest.Get().(*TaobaoLogisticsAppletPackageQueryAPIRequest)
}

// ReleaseTaobaoLogisticsAppletPackageQueryAPIRequest 将 TaobaoLogisticsAppletPackageQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsAppletPackageQueryAPIRequest(v *TaobaoLogisticsAppletPackageQueryAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsAppletPackageQueryAPIRequest.Put(v)
}
