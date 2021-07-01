package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPopupstoreActivityDeviceQueryAPIRequest
根据活动id查询活动相关快闪店及设备信息 API请求
tmall.popupstore.activity.device.query

查询某一活动的deviceCode的部署情况 */
type TmallPopupstoreActivityDeviceQueryAPIRequest struct {
	model.Params
	// ISV的活动ID
	_activityId int64
}

// NewTmallPopupstoreActivityDeviceQueryRequest 初始化TmallPopupstoreActivityDeviceQueryAPIRequest对象
func NewTmallPopupstoreActivityDeviceQueryRequest() *TmallPopupstoreActivityDeviceQueryAPIRequest {
	return &TmallPopupstoreActivityDeviceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.activity.device.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// ISV的活动ID
func (r *TmallPopupstoreActivityDeviceQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}
