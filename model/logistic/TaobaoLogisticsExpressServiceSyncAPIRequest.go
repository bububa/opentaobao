package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressservicesyncAPIRequest 服务信息回告接口 API请求
// taobao.logistics.express.service.sync
//
// 服务信息回告接口
type TaobaologisticsexpressservicesyncAPIRequest struct {
	model.Params
	// 参数信息
	_tmsServiceSyncRequest *TmsServiceSyncRequest
}

// NewTaobaologisticsexpressservicesyncRequest 初始化TaobaologisticsexpressservicesyncAPIRequest对象
func NewTaobaologisticsexpressservicesyncRequest() *TaobaologisticsexpressservicesyncAPIRequest {
	return &TaobaologisticsexpressservicesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressservicesyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.service.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressservicesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressservicesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsServiceSyncRequest is TmsServiceSyncRequest Setter
// 参数信息
func (r *TaobaologisticsexpressservicesyncAPIRequest) SetTmsServiceSyncRequest(_tmsServiceSyncRequest *TmsServiceSyncRequest) error {
	r._tmsServiceSyncRequest = _tmsServiceSyncRequest
	r.Set("tms_service_sync_request", _tmsServiceSyncRequest)
	return nil
}

// GetTmsServiceSyncRequest TmsServiceSyncRequest Getter
func (r TaobaologisticsexpressservicesyncAPIRequest) GetTmsServiceSyncRequest() *TmsServiceSyncRequest {
	return r._tmsServiceSyncRequest
}
