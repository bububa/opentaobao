package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageUpdateAPIRequest 服务商更新预警消息状态 API请求
// tmall.servicecenter.servicemonitormessage.update
//
// 服务商收到预警后，需要进行回复已读状态，并可填写备注
type TmallServicecenterServicemonitormessageUpdateAPIRequest struct {
	model.Params
	// 预警处理备注
	_memo string
	// 预警消息id
	_serviceMonitorMessageId int64
	// 可更新状态：3、已读
	_status int64
}

// NewTmallServicecenterServicemonitormessageUpdateRequest 初始化TmallServicecenterServicemonitormessageUpdateAPIRequest对象
func NewTmallServicecenterServicemonitormessageUpdateRequest() *TmallServicecenterServicemonitormessageUpdateAPIRequest {
	return &TmallServicecenterServicemonitormessageUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) Reset() {
	r._memo = ""
	r._serviceMonitorMessageId = 0
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 预警处理备注
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetServiceMonitorMessageId is ServiceMonitorMessageId Setter
// 预警消息id
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetServiceMonitorMessageId(_serviceMonitorMessageId int64) error {
	r._serviceMonitorMessageId = _serviceMonitorMessageId
	r.Set("service_monitor_message_id", _serviceMonitorMessageId)
	return nil
}

// GetServiceMonitorMessageId ServiceMonitorMessageId Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetServiceMonitorMessageId() int64 {
	return r._serviceMonitorMessageId
}

// SetStatus is Status Setter
// 可更新状态：3、已读
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

var poolTmallServicecenterServicemonitormessageUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicemonitormessageUpdateRequest()
	},
}

// GetTmallServicecenterServicemonitormessageUpdateRequest 从 sync.Pool 获取 TmallServicecenterServicemonitormessageUpdateAPIRequest
func GetTmallServicecenterServicemonitormessageUpdateAPIRequest() *TmallServicecenterServicemonitormessageUpdateAPIRequest {
	return poolTmallServicecenterServicemonitormessageUpdateAPIRequest.Get().(*TmallServicecenterServicemonitormessageUpdateAPIRequest)
}

// ReleaseTmallServicecenterServicemonitormessageUpdateAPIRequest 将 TmallServicecenterServicemonitormessageUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageUpdateAPIRequest(v *TmallServicecenterServicemonitormessageUpdateAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageUpdateAPIRequest.Put(v)
}
