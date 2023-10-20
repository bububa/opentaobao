package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkresourcegroupqueryAPIRequest 查询网格仓-区块-自提点关系 API请求
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
type WdklogisticnetworkresourcegroupqueryAPIRequest struct {
	model.Params
	// 入参
	_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest
}

// NewWdklogisticnetworkresourcegroupqueryRequest 初始化WdklogisticnetworkresourcegroupqueryAPIRequest对象
func NewWdklogisticnetworkresourcegroupqueryRequest() *WdklogisticnetworkresourcegroupqueryAPIRequest {
	return &WdklogisticnetworkresourcegroupqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdklogisticnetworkresourcegroupqueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.resource.group.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdklogisticnetworkresourcegroupqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdklogisticnetworkresourcegroupqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamResourceGroupPageQueryRequest is ParamResourceGroupPageQueryRequest Setter
// 入参
func (r *WdklogisticnetworkresourcegroupqueryAPIRequest) SetParamResourceGroupPageQueryRequest(_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest) error {
	r._paramResourceGroupPageQueryRequest = _paramResourceGroupPageQueryRequest
	r.Set("param_resource_group_page_query_request", _paramResourceGroupPageQueryRequest)
	return nil
}

// GetParamResourceGroupPageQueryRequest ParamResourceGroupPageQueryRequest Getter
func (r WdklogisticnetworkresourcegroupqueryAPIRequest) GetParamResourceGroupPageQueryRequest() *ResourceGroupPageQueryRequest {
	return r._paramResourceGroupPageQueryRequest
}
