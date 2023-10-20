package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardReservefailAPIRequest 预约失败 API请求
// tmall.servicecenter.workcard.reservefail
//
// 服务商调用该接口回传工单预约失败
type TmallServicecenterWorkcardReservefailAPIRequest struct {
	model.Params
	// 下次联系时间
	_gmtNextContact string
	// 预约失败原因描述
	_failDesc string
	// 核销单外部id
	_identifyTaskId int64
	// 预约失败原因码
	_failCode int64
	// 工单id
	_workcardId int64
}

// NewTmallServicecenterWorkcardReservefailRequest 初始化TmallServicecenterWorkcardReservefailAPIRequest对象
func NewTmallServicecenterWorkcardReservefailRequest() *TmallServicecenterWorkcardReservefailAPIRequest {
	return &TmallServicecenterWorkcardReservefailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.reservefail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtNextContact is GmtNextContact Setter
// 下次联系时间
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetGmtNextContact(_gmtNextContact string) error {
	r._gmtNextContact = _gmtNextContact
	r.Set("gmt_next_contact", _gmtNextContact)
	return nil
}

// GetGmtNextContact GmtNextContact Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetGmtNextContact() string {
	return r._gmtNextContact
}

// SetFailDesc is FailDesc Setter
// 预约失败原因描述
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetFailDesc(_failDesc string) error {
	r._failDesc = _failDesc
	r.Set("fail_desc", _failDesc)
	return nil
}

// GetFailDesc FailDesc Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetFailDesc() string {
	return r._failDesc
}

// SetIdentifyTaskId is IdentifyTaskId Setter
// 核销单外部id
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetIdentifyTaskId(_identifyTaskId int64) error {
	r._identifyTaskId = _identifyTaskId
	r.Set("identify_task_id", _identifyTaskId)
	return nil
}

// GetIdentifyTaskId IdentifyTaskId Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetIdentifyTaskId() int64 {
	return r._identifyTaskId
}

// SetFailCode is FailCode Setter
// 预约失败原因码
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetFailCode(_failCode int64) error {
	r._failCode = _failCode
	r.Set("fail_code", _failCode)
	return nil
}

// GetFailCode FailCode Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetFailCode() int64 {
	return r._failCode
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
