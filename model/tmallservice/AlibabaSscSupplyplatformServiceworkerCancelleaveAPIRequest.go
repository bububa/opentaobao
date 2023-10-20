package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest 工人取消请假 API请求
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
type AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest struct {
	model.Params
	// 时间段
	_leaveBeginAndEndList []string
	// 身份证号
	_identityNumber string
}

// NewAlibabasscsupplyplatformserviceworkercancelleaveRequest 初始化AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest对象
func NewAlibabasscsupplyplatformserviceworkercancelleaveRequest() *AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest {
	return &AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.cancelleave"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeaveBeginAndEndList is LeaveBeginAndEndList Setter
// 时间段
func (r *AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) SetLeaveBeginAndEndList(_leaveBeginAndEndList []string) error {
	r._leaveBeginAndEndList = _leaveBeginAndEndList
	r.Set("leave_begin_and_end_list", _leaveBeginAndEndList)
	return nil
}

// GetLeaveBeginAndEndList LeaveBeginAndEndList Getter
func (r AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) GetLeaveBeginAndEndList() []string {
	return r._leaveBeginAndEndList
}

// SetIdentityNumber is IdentityNumber Setter
// 身份证号
func (r *AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) SetIdentityNumber(_identityNumber string) error {
	r._identityNumber = _identityNumber
	r.Set("identity_number", _identityNumber)
	return nil
}

// GetIdentityNumber IdentityNumber Getter
func (r AlibabasscsupplyplatformserviceworkercancelleaveAPIRequest) GetIdentityNumber() string {
	return r._identityNumber
}
