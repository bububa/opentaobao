package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtSceneActivityQueryAPIRequest 喵零场景活动查询 API请求
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
type TmallNrtSceneActivityQueryAPIRequest struct {
	model.Params
	// 业务身份
	_bizCode string
	// 活动id
	_activityId int64
}

// NewTmallNrtSceneActivityQueryRequest 初始化TmallNrtSceneActivityQueryAPIRequest对象
func NewTmallNrtSceneActivityQueryRequest() *TmallNrtSceneActivityQueryAPIRequest {
	return &TmallNrtSceneActivityQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtSceneActivityQueryAPIRequest) Reset() {
	r._bizCode = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtSceneActivityQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.scene.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtSceneActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtSceneActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallNrtSceneActivityQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtSceneActivityQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TmallNrtSceneActivityQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TmallNrtSceneActivityQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTmallNrtSceneActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtSceneActivityQueryRequest()
	},
}

// GetTmallNrtSceneActivityQueryRequest 从 sync.Pool 获取 TmallNrtSceneActivityQueryAPIRequest
func GetTmallNrtSceneActivityQueryAPIRequest() *TmallNrtSceneActivityQueryAPIRequest {
	return poolTmallNrtSceneActivityQueryAPIRequest.Get().(*TmallNrtSceneActivityQueryAPIRequest)
}

// ReleaseTmallNrtSceneActivityQueryAPIRequest 将 TmallNrtSceneActivityQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtSceneActivityQueryAPIRequest(v *TmallNrtSceneActivityQueryAPIRequest) {
	v.Reset()
	poolTmallNrtSceneActivityQueryAPIRequest.Put(v)
}
