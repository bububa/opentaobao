package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardSuspendAPIRequest 工单挂起 API请求
// tmall.servicecenter.workcard.suspend
//
// 工单挂起
type TmallServicecenterWorkcardSuspendAPIRequest struct {
	model.Params
	// 预约时间
	_reserveServiceDate string
	// 下次联系时间
	_gmtNextContact string
	// 工单id
	_workcardId int64
	// 挂起原因类型code
	_failCode int64
	// 挂起原因描述
	_failDesc string
}

// NewTmallServicecenterWorkcardSuspendRequest 初始化TmallServicecenterWorkcardSuspendAPIRequest对象
func NewTmallServicecenterWorkcardSuspendRequest() *TmallServicecenterWorkcardSuspendAPIRequest {
	return &TmallServicecenterWorkcardSuspendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.suspend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReserveServiceDate Setter
// 预约时间
func (r *TmallServicecenterWorkcardSuspendAPIRequest) SetReserveServiceDate(_reserveServiceDate string) error {
	r._reserveServiceDate = _reserveServiceDate
	r.Set("reserve_service_date", _reserveServiceDate)
	return nil
}

// Get ReserveServiceDate Getter
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetReserveServiceDate() string {
	return r._reserveServiceDate
}

// Set is GmtNextContact Setter
// 下次联系时间
func (r *TmallServicecenterWorkcardSuspendAPIRequest) SetGmtNextContact(_gmtNextContact string) error {
	r._gmtNextContact = _gmtNextContact
	r.Set("gmt_next_contact", _gmtNextContact)
	return nil
}

// Get GmtNextContact Getter
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetGmtNextContact() string {
	return r._gmtNextContact
}

// Set is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardSuspendAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// Get WorkcardId Getter
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// Set is FailCode Setter
// 挂起原因类型code
func (r *TmallServicecenterWorkcardSuspendAPIRequest) SetFailCode(_failCode int64) error {
	r._failCode = _failCode
	r.Set("fail_code", _failCode)
	return nil
}

// Get FailCode Getter
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetFailCode() int64 {
	return r._failCode
}

// Set is FailDesc Setter
// 挂起原因描述
func (r *TmallServicecenterWorkcardSuspendAPIRequest) SetFailDesc(_failDesc string) error {
	r._failDesc = _failDesc
	r.Set("fail_desc", _failDesc)
	return nil
}

// Get FailDesc Getter
func (r TmallServicecenterWorkcardSuspendAPIRequest) GetFailDesc() string {
	return r._failDesc
}
