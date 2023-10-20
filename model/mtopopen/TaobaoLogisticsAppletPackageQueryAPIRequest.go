package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsappletpackagequeryAPIRequest 淘宝包裹查询 API请求
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
type TaobaologisticsappletpackagequeryAPIRequest struct {
	model.Params
	// 查询用户的包裹列表
	_queryPackageListRequest *QueryPackageListRequest
}

// NewTaobaologisticsappletpackagequeryRequest 初始化TaobaologisticsappletpackagequeryAPIRequest对象
func NewTaobaologisticsappletpackagequeryRequest() *TaobaologisticsappletpackagequeryAPIRequest {
	return &TaobaologisticsappletpackagequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsappletpackagequeryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.applet.package.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsappletpackagequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsappletpackagequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryPackageListRequest is QueryPackageListRequest Setter
// 查询用户的包裹列表
func (r *TaobaologisticsappletpackagequeryAPIRequest) SetQueryPackageListRequest(_queryPackageListRequest *QueryPackageListRequest) error {
	r._queryPackageListRequest = _queryPackageListRequest
	r.Set("query_package_list_request", _queryPackageListRequest)
	return nil
}

// GetQueryPackageListRequest QueryPackageListRequest Getter
func (r TaobaologisticsappletpackagequeryAPIRequest) GetQueryPackageListRequest() *QueryPackageListRequest {
	return r._queryPackageListRequest
}
