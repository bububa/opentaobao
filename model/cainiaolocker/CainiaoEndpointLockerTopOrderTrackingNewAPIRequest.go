package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertopordertrackingnewAPIRequest 事件回传接口 API请求
// cainiao.endpoint.locker.top.order.tracking.new
//
// 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoendpointlockertopordertrackingnewAPIRequest struct {
	model.Params
	// 回传信息
	_trackInfo *CollectTrackingInfo
}

// NewCainiaoendpointlockertopordertrackingnewRequest 初始化CainiaoendpointlockertopordertrackingnewAPIRequest对象
func NewCainiaoendpointlockertopordertrackingnewRequest() *CainiaoendpointlockertopordertrackingnewAPIRequest {
	return &CainiaoendpointlockertopordertrackingnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoendpointlockertopordertrackingnewAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoendpointlockertopordertrackingnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoendpointlockertopordertrackingnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackInfo is TrackInfo Setter
// 回传信息
func (r *CainiaoendpointlockertopordertrackingnewAPIRequest) SetTrackInfo(_trackInfo *CollectTrackingInfo) error {
	r._trackInfo = _trackInfo
	r.Set("track_info", _trackInfo)
	return nil
}

// GetTrackInfo TrackInfo Getter
func (r CainiaoendpointlockertopordertrackingnewAPIRequest) GetTrackInfo() *CollectTrackingInfo {
	return r._trackInfo
}
