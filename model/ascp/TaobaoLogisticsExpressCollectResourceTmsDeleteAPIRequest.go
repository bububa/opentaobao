package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest 上门取退可揽范围删除 API请求
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
type TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest struct {
	model.Params
	// 上门取退可揽范围删除
	_collectResourceDeleteRequest *CollectResourceDeleteRequest
}

// NewTaobaoLogisticsExpressCollectResourceTmsDeleteRequest 初始化TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest对象
func NewTaobaoLogisticsExpressCollectResourceTmsDeleteRequest() *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest {
	return &TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) Reset() {
	r._collectResourceDeleteRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.collect.resource.tms.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCollectResourceDeleteRequest is CollectResourceDeleteRequest Setter
// 上门取退可揽范围删除
func (r *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) SetCollectResourceDeleteRequest(_collectResourceDeleteRequest *CollectResourceDeleteRequest) error {
	r._collectResourceDeleteRequest = _collectResourceDeleteRequest
	r.Set("collect_resource_delete_request", _collectResourceDeleteRequest)
	return nil
}

// GetCollectResourceDeleteRequest CollectResourceDeleteRequest Getter
func (r TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) GetCollectResourceDeleteRequest() *CollectResourceDeleteRequest {
	return r._collectResourceDeleteRequest
}

var poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressCollectResourceTmsDeleteRequest()
	},
}

// GetTaobaoLogisticsExpressCollectResourceTmsDeleteRequest 从 sync.Pool 获取 TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest
func GetTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest() *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest {
	return poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest.Get().(*TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest)
}

// ReleaseTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest 将 TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest(v *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIRequest.Put(v)
}
