package nrt

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtSceneActivityQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.scene.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtSceneActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
