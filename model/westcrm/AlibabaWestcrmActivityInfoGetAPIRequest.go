package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmActivityInfoGetAPIRequest 获取活动详情 API请求
// alibaba.westcrm.activity.info.get
//
// 根据id查询活动详情
type AlibabaWestcrmActivityInfoGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 活动id
	_activityId int64
}

// NewAlibabaWestcrmActivityInfoGetRequest 初始化AlibabaWestcrmActivityInfoGetAPIRequest对象
func NewAlibabaWestcrmActivityInfoGetRequest() *AlibabaWestcrmActivityInfoGetAPIRequest {
	return &AlibabaWestcrmActivityInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmActivityInfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.activity.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmActivityInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaWestcrmActivityInfoGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaWestcrmActivityInfoGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaWestcrmActivityInfoGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaWestcrmActivityInfoGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
