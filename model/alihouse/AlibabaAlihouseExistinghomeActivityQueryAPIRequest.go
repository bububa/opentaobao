package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeactivityqueryAPIRequest 五福活动经纪人获取 API请求
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
type AlibabaalihouseexistinghomeactivityqueryAPIRequest struct {
	model.Params
	// 查询的参数
	_queryInfo *ActivityQueryInfoDto
}

// NewAlibabaalihouseexistinghomeactivityqueryRequest 初始化AlibabaalihouseexistinghomeactivityqueryAPIRequest对象
func NewAlibabaalihouseexistinghomeactivityqueryRequest() *AlibabaalihouseexistinghomeactivityqueryAPIRequest {
	return &AlibabaalihouseexistinghomeactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryInfo is QueryInfo Setter
// 查询的参数
func (r *AlibabaalihouseexistinghomeactivityqueryAPIRequest) SetQueryInfo(_queryInfo *ActivityQueryInfoDto) error {
	r._queryInfo = _queryInfo
	r.Set("query_info", _queryInfo)
	return nil
}

// GetQueryInfo QueryInfo Getter
func (r AlibabaalihouseexistinghomeactivityqueryAPIRequest) GetQueryInfo() *ActivityQueryInfoDto {
	return r._queryInfo
}
