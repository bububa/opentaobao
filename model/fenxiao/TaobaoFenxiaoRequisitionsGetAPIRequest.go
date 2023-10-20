package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaorequisitionsgetAPIRequest 合作申请查询 API请求
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
type TaobaofenxiaorequisitionsgetAPIRequest struct {
	model.Params
	// 申请开始时间yyyy-MM-dd
	_applyStart string
	// 申请结束时间yyyy-MM-dd
	_applyEnd string
	// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
	_status int64
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
}

// NewTaobaofenxiaorequisitionsgetRequest 初始化TaobaofenxiaorequisitionsgetAPIRequest对象
func NewTaobaofenxiaorequisitionsgetRequest() *TaobaofenxiaorequisitionsgetAPIRequest {
	return &TaobaofenxiaorequisitionsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.requisitions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyStart is ApplyStart Setter
// 申请开始时间yyyy-MM-dd
func (r *TaobaofenxiaorequisitionsgetAPIRequest) SetApplyStart(_applyStart string) error {
	r._applyStart = _applyStart
	r.Set("apply_start", _applyStart)
	return nil
}

// GetApplyStart ApplyStart Getter
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetApplyStart() string {
	return r._applyStart
}

// SetApplyEnd is ApplyEnd Setter
// 申请结束时间yyyy-MM-dd
func (r *TaobaofenxiaorequisitionsgetAPIRequest) SetApplyEnd(_applyEnd string) error {
	r._applyEnd = _applyEnd
	r.Set("apply_end", _applyEnd)
	return nil
}

// GetApplyEnd ApplyEnd Getter
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetApplyEnd() string {
	return r._applyEnd
}

// SetStatus is Status Setter
// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
func (r *TaobaofenxiaorequisitionsgetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaofenxiaorequisitionsgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaofenxiaorequisitionsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofenxiaorequisitionsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
