package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkLogisticNetworkResourceGroupQueryAPIRequest) Reset() {
	r._paramResourceGroupPageQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetApiMethodName() string {
	return "wdk.logistic.network.resource.group.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkLogisticNetworkResourceGroupQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolWdkLogisticNetworkResourceGroupQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkLogisticNetworkResourceGroupQueryRequest()
	},
}

// GetWdkLogisticNetworkResourceGroupQueryRequest 从 sync.Pool 获取 WdkLogisticNetworkResourceGroupQueryAPIRequest
func GetWdkLogisticNetworkResourceGroupQueryAPIRequest() *WdkLogisticNetworkResourceGroupQueryAPIRequest {
	return poolWdkLogisticNetworkResourceGroupQueryAPIRequest.Get().(*WdkLogisticNetworkResourceGroupQueryAPIRequest)
}

// ReleaseWdkLogisticNetworkResourceGroupQueryAPIRequest 将 WdkLogisticNetworkResourceGroupQueryAPIRequest 放入 sync.Pool
func ReleaseWdkLogisticNetworkResourceGroupQueryAPIRequest(v *WdkLogisticNetworkResourceGroupQueryAPIRequest) {
	v.Reset()
	poolWdkLogisticNetworkResourceGroupQueryAPIRequest.Put(v)
}
