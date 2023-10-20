package alitrippoi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformPoiRawSaverawpoiAPIRequest POI开放存储能力 API请求
// alitrip.platform.poi.raw.saverawpoi
//
// POI开放存储提供离线/在线/纬错更新的能力
type AlitripPlatformPoiRawSaverawpoiAPIRequest struct {
	model.Params
	// poi存储参数
	_tripPoiRawSaveParam *TripPoiRawSaveParamV2
}

// NewAlitripPlatformPoiRawSaverawpoiRequest 初始化AlitripPlatformPoiRawSaverawpoiAPIRequest对象
func NewAlitripPlatformPoiRawSaverawpoiRequest() *AlitripPlatformPoiRawSaverawpoiAPIRequest {
	return &AlitripPlatformPoiRawSaverawpoiAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformPoiRawSaverawpoiAPIRequest) Reset() {
	r._tripPoiRawSaveParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.poi.raw.saverawpoi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTripPoiRawSaveParam is TripPoiRawSaveParam Setter
// poi存储参数
func (r *AlitripPlatformPoiRawSaverawpoiAPIRequest) SetTripPoiRawSaveParam(_tripPoiRawSaveParam *TripPoiRawSaveParamV2) error {
	r._tripPoiRawSaveParam = _tripPoiRawSaveParam
	r.Set("trip_poi_raw_save_param", _tripPoiRawSaveParam)
	return nil
}

// GetTripPoiRawSaveParam TripPoiRawSaveParam Getter
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetTripPoiRawSaveParam() *TripPoiRawSaveParamV2 {
	return r._tripPoiRawSaveParam
}

var poolAlitripPlatformPoiRawSaverawpoiAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformPoiRawSaverawpoiRequest()
	},
}

// GetAlitripPlatformPoiRawSaverawpoiRequest 从 sync.Pool 获取 AlitripPlatformPoiRawSaverawpoiAPIRequest
func GetAlitripPlatformPoiRawSaverawpoiAPIRequest() *AlitripPlatformPoiRawSaverawpoiAPIRequest {
	return poolAlitripPlatformPoiRawSaverawpoiAPIRequest.Get().(*AlitripPlatformPoiRawSaverawpoiAPIRequest)
}

// ReleaseAlitripPlatformPoiRawSaverawpoiAPIRequest 将 AlitripPlatformPoiRawSaverawpoiAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformPoiRawSaverawpoiAPIRequest(v *AlitripPlatformPoiRawSaverawpoiAPIRequest) {
	v.Reset()
	poolAlitripPlatformPoiRawSaverawpoiAPIRequest.Put(v)
}
