package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreActivityQueryAPIRequest 查询某段时间内的快闪活动列表 API请求
// tmall.popupstore.activity.query
//
// 提供给ISV查询某一时间段内包含指定appKey的活动列表
type TmallPopupstoreActivityQueryAPIRequest struct {
	model.Params
	// 查询开始时间,yyyy-MM-dd
	_startDate string
	// 查询结束时间，yyyy-MM-dd
	_endDate string
}

// NewTmallPopupstoreActivityQueryRequest 初始化TmallPopupstoreActivityQueryAPIRequest对象
func NewTmallPopupstoreActivityQueryRequest() *TmallPopupstoreActivityQueryAPIRequest {
	return &TmallPopupstoreActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPopupstoreActivityQueryAPIRequest) GetApiMethodName() string {
	return "tmall.popupstore.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPopupstoreActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPopupstoreActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 查询开始时间,yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TmallPopupstoreActivityQueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询结束时间，yyyy-MM-dd
func (r *TmallPopupstoreActivityQueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TmallPopupstoreActivityQueryAPIRequest) GetEndDate() string {
	return r._endDate
}
