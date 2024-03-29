package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressServiceSyncAPIRequest 服务信息回告接口 API请求
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
type TaobaoLogisticsExpressServiceSyncAPIRequest struct {
	model.Params
	// 参数信息
	_tmsServiceSyncRequest *TmsServiceSyncRequest
}

// NewTaobaoLogisticsExpressServiceSyncRequest 初始化TaobaoLogisticsExpressServiceSyncAPIRequest对象
func NewTaobaoLogisticsExpressServiceSyncRequest() *TaobaoLogisticsExpressServiceSyncAPIRequest {
	return &TaobaoLogisticsExpressServiceSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressServiceSyncAPIRequest) Reset() {
	r._tmsServiceSyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressServiceSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.service.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressServiceSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressServiceSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsServiceSyncRequest is TmsServiceSyncRequest Setter
// 参数信息
func (r *TaobaoLogisticsExpressServiceSyncAPIRequest) SetTmsServiceSyncRequest(_tmsServiceSyncRequest *TmsServiceSyncRequest) error {
	r._tmsServiceSyncRequest = _tmsServiceSyncRequest
	r.Set("tms_service_sync_request", _tmsServiceSyncRequest)
	return nil
}

// GetTmsServiceSyncRequest TmsServiceSyncRequest Getter
func (r TaobaoLogisticsExpressServiceSyncAPIRequest) GetTmsServiceSyncRequest() *TmsServiceSyncRequest {
	return r._tmsServiceSyncRequest
}

var poolTaobaoLogisticsExpressServiceSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressServiceSyncRequest()
	},
}

// GetTaobaoLogisticsExpressServiceSyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressServiceSyncAPIRequest
func GetTaobaoLogisticsExpressServiceSyncAPIRequest() *TaobaoLogisticsExpressServiceSyncAPIRequest {
	return poolTaobaoLogisticsExpressServiceSyncAPIRequest.Get().(*TaobaoLogisticsExpressServiceSyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressServiceSyncAPIRequest 将 TaobaoLogisticsExpressServiceSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressServiceSyncAPIRequest(v *TaobaoLogisticsExpressServiceSyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressServiceSyncAPIRequest.Put(v)
}
