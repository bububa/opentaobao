package smartstore

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreActivityDeviceQueryAPIRequest 根据活动id查询活动相关快闪店及设备信息 API请求
// tmall.popupstore.activity.device.query
//
// 查询某一活动的deviceCode的部署情况
type TmallPopupstoreActivityDeviceQueryAPIRequest struct {
	model.Params
	// ISV的活动ID
	_activityId int64
}

// NewTmallPopupstoreActivityDeviceQueryRequest 初始化TmallPopupstoreActivityDeviceQueryAPIRequest对象
func NewTmallPopupstoreActivityDeviceQueryRequest() *TmallPopupstoreActivityDeviceQueryAPIRequest {
	return &TmallPopupstoreActivityDeviceQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPopupstoreActivityDeviceQueryAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.activity.device.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// ISV的活动ID
func (r *TmallPopupstoreActivityDeviceQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TmallPopupstoreActivityDeviceQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTmallPopupstoreActivityDeviceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPopupstoreActivityDeviceQueryRequest()
	},
}

// GetTmallPopupstoreActivityDeviceQueryRequest 从 sync.Pool 获取 TmallPopupstoreActivityDeviceQueryAPIRequest
func GetTmallPopupstoreActivityDeviceQueryAPIRequest() *TmallPopupstoreActivityDeviceQueryAPIRequest {
	return poolTmallPopupstoreActivityDeviceQueryAPIRequest.Get().(*TmallPopupstoreActivityDeviceQueryAPIRequest)
}

// ReleaseTmallPopupstoreActivityDeviceQueryAPIRequest 将 TmallPopupstoreActivityDeviceQueryAPIRequest 放入 sync.Pool
func ReleaseTmallPopupstoreActivityDeviceQueryAPIRequest(v *TmallPopupstoreActivityDeviceQueryAPIRequest) {
	v.Reset()
	poolTmallPopupstoreActivityDeviceQueryAPIRequest.Put(v)
}
