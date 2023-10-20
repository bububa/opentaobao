package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicemonitormessageupdateAPIRequest 服务商更新预警消息状态 API请求
// tmall.servicecenter.servicemonitormessage.update
//
// 服务商收到预警后，需要进行回复已读状态，并可填写备注
type TmallservicecenterservicemonitormessageupdateAPIRequest struct {
	model.Params
	// 预警处理备注
	_memo string
	// 预警消息id
	_serviceMonitorMessageId int64
	// 可更新状态：3、已读
	_status int64
}

// NewTmallservicecenterservicemonitormessageupdateRequest 初始化TmallservicecenterservicemonitormessageupdateAPIRequest对象
func NewTmallservicecenterservicemonitormessageupdateRequest() *TmallservicecenterservicemonitormessageupdateAPIRequest {
	return &TmallservicecenterservicemonitormessageupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 预警处理备注
func (r *TmallservicecenterservicemonitormessageupdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetServiceMonitorMessageId is ServiceMonitorMessageId Setter
// 预警消息id
func (r *TmallservicecenterservicemonitormessageupdateAPIRequest) SetServiceMonitorMessageId(_serviceMonitorMessageId int64) error {
	r._serviceMonitorMessageId = _serviceMonitorMessageId
	r.Set("service_monitor_message_id", _serviceMonitorMessageId)
	return nil
}

// GetServiceMonitorMessageId ServiceMonitorMessageId Getter
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetServiceMonitorMessageId() int64 {
	return r._serviceMonitorMessageId
}

// SetStatus is Status Setter
// 可更新状态：3、已读
func (r *TmallservicecenterservicemonitormessageupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallservicecenterservicemonitormessageupdateAPIRequest) GetStatus() int64 {
	return r._status
}
