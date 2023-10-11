package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest 配服务商揽收能力同步接口 API请求
// taobao.logistics.express.collect.resource.tms.async
//
// 配服务商揽收能力同步接口
type TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest struct {
	model.Params
	// 上门取退可揽范围
	_collectResourceRequest *CollectResourceRequest
}

// NewTaobaoLogisticsExpressCollectResourceTmsAsyncRequest 初始化TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest对象
func NewTaobaoLogisticsExpressCollectResourceTmsAsyncRequest() *TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest {
	return &TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.collect.resource.tms.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCollectResourceRequest is CollectResourceRequest Setter
// 上门取退可揽范围
func (r *TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest) SetCollectResourceRequest(_collectResourceRequest *CollectResourceRequest) error {
	r._collectResourceRequest = _collectResourceRequest
	r.Set("collect_resource_request", _collectResourceRequest)
	return nil
}

// GetCollectResourceRequest CollectResourceRequest Getter
func (r TaobaoLogisticsExpressCollectResourceTmsAsyncAPIRequest) GetCollectResourceRequest() *CollectResourceRequest {
	return r._collectResourceRequest
}
