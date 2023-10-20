package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpopupstoreactivitydevicequeryAPIRequest 根据活动id查询活动相关快闪店及设备信息 API请求
// tmall.popupstore.activity.device.query
//
// 查询某一活动的deviceCode的部署情况
type TmallpopupstoreactivitydevicequeryAPIRequest struct {
	model.Params
	// ISV的活动ID
	_activityId int64
}

// NewTmallpopupstoreactivitydevicequeryRequest 初始化TmallpopupstoreactivitydevicequeryAPIRequest对象
func NewTmallpopupstoreactivitydevicequeryRequest() *TmallpopupstoreactivitydevicequeryAPIRequest {
	return &TmallpopupstoreactivitydevicequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpopupstoreactivitydevicequeryAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.activity.device.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpopupstoreactivitydevicequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpopupstoreactivitydevicequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// ISV的活动ID
func (r *TmallpopupstoreactivitydevicequeryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TmallpopupstoreactivitydevicequeryAPIRequest) GetActivityId() int64 {
	return r._activityId
}
