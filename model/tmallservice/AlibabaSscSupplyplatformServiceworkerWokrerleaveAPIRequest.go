package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest 工人请假 API请求
// alibaba.ssc.supplyplatform.serviceworker.wokrerleave
//
// 工人请假
type AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest struct {
	model.Params
	// 开始和结束时间
	_leaveBeginAndEndArray []string
	// 请假原因
	_reason string
	// 身份证号
	_identityNumber string
}

// NewAlibabaSscSupplyplatformServiceworkerWokrerleaveRequest 初始化AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerWokrerleaveRequest() *AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.wokrerleave"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLeaveBeginAndEndArray is LeaveBeginAndEndArray Setter
// 开始和结束时间
func (r *AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) SetLeaveBeginAndEndArray(_leaveBeginAndEndArray []string) error {
	r._leaveBeginAndEndArray = _leaveBeginAndEndArray
	r.Set("leave_begin_and_end_array", _leaveBeginAndEndArray)
	return nil
}

// GetLeaveBeginAndEndArray LeaveBeginAndEndArray Getter
func (r AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) GetLeaveBeginAndEndArray() []string {
	return r._leaveBeginAndEndArray
}

// SetReason is Reason Setter
// 请假原因
func (r *AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) GetReason() string {
	return r._reason
}

// SetIdentityNumber is IdentityNumber Setter
// 身份证号
func (r *AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) SetIdentityNumber(_identityNumber string) error {
	r._identityNumber = _identityNumber
	r.Set("identity_number", _identityNumber)
	return nil
}

// GetIdentityNumber IdentityNumber Getter
func (r AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest) GetIdentityNumber() string {
	return r._identityNumber
}
