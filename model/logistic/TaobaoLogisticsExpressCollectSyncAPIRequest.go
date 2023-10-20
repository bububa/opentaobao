package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscollectsyncAPIRequest 服饰逆向揽收信息同步 API请求
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
type TaobaologisticsexpresscollectsyncAPIRequest struct {
	model.Params
	// 揽收参数
	_tmsCollectRequest *TmsCollectRequest
}

// NewTaobaologisticsexpresscollectsyncRequest 初始化TaobaologisticsexpresscollectsyncAPIRequest对象
func NewTaobaologisticsexpresscollectsyncRequest() *TaobaologisticsexpresscollectsyncAPIRequest {
	return &TaobaologisticsexpresscollectsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresscollectsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.collect.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresscollectsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresscollectsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCollectRequest is TmsCollectRequest Setter
// 揽收参数
func (r *TaobaologisticsexpresscollectsyncAPIRequest) SetTmsCollectRequest(_tmsCollectRequest *TmsCollectRequest) error {
	r._tmsCollectRequest = _tmsCollectRequest
	r.Set("tms_collect_request", _tmsCollectRequest)
	return nil
}

// GetTmsCollectRequest TmsCollectRequest Getter
func (r TaobaologisticsexpresscollectsyncAPIRequest) GetTmsCollectRequest() *TmsCollectRequest {
	return r._tmsCollectRequest
}
