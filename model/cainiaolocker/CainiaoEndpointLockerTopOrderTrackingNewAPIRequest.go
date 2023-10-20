package cainiaolocker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopOrderTrackingNewAPIRequest 事件回传接口 API请求
// cainiao.endpoint.locker.top.order.tracking.new
//
// 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
type CainiaoEndpointLockerTopOrderTrackingNewAPIRequest struct {
	model.Params
	// 回传信息
	_trackInfo *CollectTrackingInfo
}

// NewCainiaoEndpointLockerTopOrderTrackingNewRequest 初始化CainiaoEndpointLockerTopOrderTrackingNewAPIRequest对象
func NewCainiaoEndpointLockerTopOrderTrackingNewRequest() *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest {
	return &CainiaoEndpointLockerTopOrderTrackingNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) Reset() {
	r._trackInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackInfo is TrackInfo Setter
// 回传信息
func (r *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) SetTrackInfo(_trackInfo *CollectTrackingInfo) error {
	r._trackInfo = _trackInfo
	r.Set("track_info", _trackInfo)
	return nil
}

// GetTrackInfo TrackInfo Getter
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetTrackInfo() *CollectTrackingInfo {
	return r._trackInfo
}

var poolCainiaoEndpointLockerTopOrderTrackingNewAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoEndpointLockerTopOrderTrackingNewRequest()
	},
}

// GetCainiaoEndpointLockerTopOrderTrackingNewRequest 从 sync.Pool 获取 CainiaoEndpointLockerTopOrderTrackingNewAPIRequest
func GetCainiaoEndpointLockerTopOrderTrackingNewAPIRequest() *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest {
	return poolCainiaoEndpointLockerTopOrderTrackingNewAPIRequest.Get().(*CainiaoEndpointLockerTopOrderTrackingNewAPIRequest)
}

// ReleaseCainiaoEndpointLockerTopOrderTrackingNewAPIRequest 将 CainiaoEndpointLockerTopOrderTrackingNewAPIRequest 放入 sync.Pool
func ReleaseCainiaoEndpointLockerTopOrderTrackingNewAPIRequest(v *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) {
	v.Reset()
	poolCainiaoEndpointLockerTopOrderTrackingNewAPIRequest.Put(v)
}
