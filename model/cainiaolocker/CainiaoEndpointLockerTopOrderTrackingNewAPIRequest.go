package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderTrackingNewAPIRequest
事件回传接口 API请求
cainiao.endpoint.locker.top.order.tracking.new

用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。 */
type CainiaoEndpointLockerTopOrderTrackingNewAPIRequest struct {
	model.Params
	// 回传信息
	_trackInfo *CollectTrackingInfo
}

// NewCainiaoEndpointLockerTopOrderTrackingNewRequest 初始化CainiaoEndpointLockerTopOrderTrackingNewAPIRequest对象
func NewCainiaoEndpointLockerTopOrderTrackingNewRequest() *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest {
	return &CainiaoEndpointLockerTopOrderTrackingNewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.tracking.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TrackInfo Setter
// 回传信息
func (r *CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) SetTrackInfo(_trackInfo *CollectTrackingInfo) error {
	r._trackInfo = _trackInfo
	r.Set("track_info", _trackInfo)
	return nil
}

// Get TrackInfo Getter
func (r CainiaoEndpointLockerTopOrderTrackingNewAPIRequest) GetTrackInfo() *CollectTrackingInfo {
	return r._trackInfo
}
