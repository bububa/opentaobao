package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkResourceGroupQueryAPIRequest 查询网格仓-区块-自提点关系 API请求
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
type WdkLogisticNetworkResourceGroupQueryAPIRequest struct {
	model.Params
	// 入参
	_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest
}

// NewWdkLogisticNetworkResourceGroupQueryRequest 初始化WdkLogisticNetworkResourceGroupQueryAPIRequest对象
func NewWdkLogisticNetworkResourceGroupQueryRequest() *WdkLogisticNetworkResourceGroupQueryAPIRequest {
	return &WdkLogisticNetworkResourceGroupQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.resource.group.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamResourceGroupPageQueryRequest is ParamResourceGroupPageQueryRequest Setter
// 入参
func (r *WdkLogisticNetworkResourceGroupQueryAPIRequest) SetParamResourceGroupPageQueryRequest(_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest) error {
	r._paramResourceGroupPageQueryRequest = _paramResourceGroupPageQueryRequest
	r.Set("param_resource_group_page_query_request", _paramResourceGroupPageQueryRequest)
	return nil
}

// GetParamResourceGroupPageQueryRequest ParamResourceGroupPageQueryRequest Getter
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetParamResourceGroupPageQueryRequest() *ResourceGroupPageQueryRequest {
	return r._paramResourceGroupPageQueryRequest
}
