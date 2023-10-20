package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitemqueryAPIRequest 查询活动下的优惠信息 API请求
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
type TaobaosingletreasureactivityitemqueryAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaosingletreasureactivityitemqueryRequest 初始化TaobaosingletreasureactivityitemqueryAPIRequest对象
func NewTaobaosingletreasureactivityitemqueryRequest() *TaobaosingletreasureactivityitemqueryAPIRequest {
	return &TaobaosingletreasureactivityitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动Id
func (r *TaobaosingletreasureactivityitemqueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaosingletreasureactivityitemqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaosingletreasureactivityitemqueryAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaosingletreasureactivityitemqueryAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
