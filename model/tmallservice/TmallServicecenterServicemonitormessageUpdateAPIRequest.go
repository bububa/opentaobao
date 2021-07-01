package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicemonitormessageUpdateAPIRequest
服务商更新预警消息状态 API请求
tmall.servicecenter.servicemonitormessage.update

服务商收到预警后，需要进行回复已读状态，并可填写备注 */
type TmallServicecenterServicemonitormessageUpdateAPIRequest struct {
	model.Params
	// 预警消息id
	_serviceMonitorMessageId int64
	// 预警处理备注
	_memo string
	// 可更新状态：3、已读
	_status int64
}

// NewTmallServicecenterServicemonitormessageUpdateRequest 初始化TmallServicecenterServicemonitormessageUpdateAPIRequest对象
func NewTmallServicecenterServicemonitormessageUpdateRequest() *TmallServicecenterServicemonitormessageUpdateAPIRequest {
	return &TmallServicecenterServicemonitormessageUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceMonitorMessageId Setter
// 预警消息id
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetServiceMonitorMessageId(_serviceMonitorMessageId int64) error {
	r._serviceMonitorMessageId = _serviceMonitorMessageId
	r.Set("service_monitor_message_id", _serviceMonitorMessageId)
	return nil
}

// Get ServiceMonitorMessageId Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetServiceMonitorMessageId() int64 {
	return r._serviceMonitorMessageId
}

// Set is Memo Setter
// 预警处理备注
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// Get Memo Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetMemo() string {
	return r._memo
}

// Set is Status Setter
// 可更新状态：3、已读
func (r *TmallServicecenterServicemonitormessageUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TmallServicecenterServicemonitormessageUpdateAPIRequest) GetStatus() int64 {
	return r._status
}
