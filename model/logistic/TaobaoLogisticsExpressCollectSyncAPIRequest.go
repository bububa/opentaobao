package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectSyncAPIRequest 服饰逆向揽收信息同步 API请求
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
type TaobaoLogisticsExpressCollectSyncAPIRequest struct {
	model.Params
	// 揽收参数
	_tmsCollectRequest *TmsCollectRequest
}

// NewTaobaoLogisticsExpressCollectSyncRequest 初始化TaobaoLogisticsExpressCollectSyncAPIRequest对象
func NewTaobaoLogisticsExpressCollectSyncRequest() *TaobaoLogisticsExpressCollectSyncAPIRequest {
	return &TaobaoLogisticsExpressCollectSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressCollectSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.collect.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressCollectSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressCollectSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCollectRequest is TmsCollectRequest Setter
// 揽收参数
func (r *TaobaoLogisticsExpressCollectSyncAPIRequest) SetTmsCollectRequest(_tmsCollectRequest *TmsCollectRequest) error {
	r._tmsCollectRequest = _tmsCollectRequest
	r.Set("tms_collect_request", _tmsCollectRequest)
	return nil
}

// GetTmsCollectRequest TmsCollectRequest Getter
func (r TaobaoLogisticsExpressCollectSyncAPIRequest) GetTmsCollectRequest() *TmsCollectRequest {
	return r._tmsCollectRequest
}
