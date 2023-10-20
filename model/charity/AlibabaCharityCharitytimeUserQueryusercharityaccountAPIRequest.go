package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest 查询用户公益账户 API请求
// alibaba.charity.charitytime.user.queryusercharityaccount
//
// 查询用户公益账户
type AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest struct {
	model.Params
	// 请求对象
	_queryCharityTimeTopApiHsfRequest *QueryCharityTimeTopApiHsfRequest
}

// NewAlibabacharitycharitytimeuserqueryusercharityaccountRequest 初始化AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest对象
func NewAlibabacharitycharitytimeuserqueryusercharityaccountRequest() *AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest {
	return &AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.user.queryusercharityaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryCharityTimeTopApiHsfRequest is QueryCharityTimeTopApiHsfRequest Setter
// 请求对象
func (r *AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest) SetQueryCharityTimeTopApiHsfRequest(_queryCharityTimeTopApiHsfRequest *QueryCharityTimeTopApiHsfRequest) error {
	r._queryCharityTimeTopApiHsfRequest = _queryCharityTimeTopApiHsfRequest
	r.Set("query_charity_time_top_api_hsf_request", _queryCharityTimeTopApiHsfRequest)
	return nil
}

// GetQueryCharityTimeTopApiHsfRequest QueryCharityTimeTopApiHsfRequest Getter
func (r AlibabacharitycharitytimeuserqueryusercharityaccountAPIRequest) GetQueryCharityTimeTopApiHsfRequest() *QueryCharityTimeTopApiHsfRequest {
	return r._queryCharityTimeTopApiHsfRequest
}
