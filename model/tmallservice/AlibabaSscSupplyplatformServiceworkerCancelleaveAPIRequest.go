package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest 工人取消请假 API请求
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest struct {
	model.Params
	// 时间段
	_leaveBeginAndEndList []string
	// 身份证号
	_identityNumber string
}

// NewAlibabaSscSupplyplatformServiceworkerCancelleaveRequest 初始化AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerCancelleaveRequest() *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) Reset() {
	r._leaveBeginAndEndList = r._leaveBeginAndEndList[:0]
	r._identityNumber = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.cancelleave"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeaveBeginAndEndList is LeaveBeginAndEndList Setter
// 时间段
func (r *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) SetLeaveBeginAndEndList(_leaveBeginAndEndList []string) error {
	r._leaveBeginAndEndList = _leaveBeginAndEndList
	r.Set("leave_begin_and_end_list", _leaveBeginAndEndList)
	return nil
}

// GetLeaveBeginAndEndList LeaveBeginAndEndList Getter
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetLeaveBeginAndEndList() []string {
	return r._leaveBeginAndEndList
}

// SetIdentityNumber is IdentityNumber Setter
// 身份证号
func (r *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) SetIdentityNumber(_identityNumber string) error {
	r._identityNumber = _identityNumber
	r.Set("identity_number", _identityNumber)
	return nil
}

// GetIdentityNumber IdentityNumber Getter
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetIdentityNumber() string {
	return r._identityNumber
}

var poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServiceworkerCancelleaveRequest()
	},
}

// GetAlibabaSscSupplyplatformServiceworkerCancelleaveRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest
func GetAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest() *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest {
	return poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest.Get().(*AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest 将 AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest(v *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest.Put(v)
}
