package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtsceneactivityqueryAPIRequest 喵零场景活动查询 API请求
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
type TmallnrtsceneactivityqueryAPIRequest struct {
	model.Params
	// 业务身份
	_bizCode string
	// 活动id
	_activityId int64
}

// NewTmallnrtsceneactivityqueryRequest 初始化TmallnrtsceneactivityqueryAPIRequest对象
func NewTmallnrtsceneactivityqueryRequest() *TmallnrtsceneactivityqueryAPIRequest {
	return &TmallnrtsceneactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtsceneactivityqueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.scene.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtsceneactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtsceneactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallnrtsceneactivityqueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtsceneactivityqueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TmallnrtsceneactivityqueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TmallnrtsceneactivityqueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}
