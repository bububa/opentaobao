package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbabyremindbatchsendAPIRequest 批量发送母婴提醒 API请求
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
type AlibabaalihealthbabyremindbatchsendAPIRequest struct {
	model.Params
	// 批量发送提醒
	_batchRemindRequests []BatchRemindRequestDto
	// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
	_remindType int64
}

// NewAlibabaalihealthbabyremindbatchsendRequest 初始化AlibabaalihealthbabyremindbatchsendAPIRequest对象
func NewAlibabaalihealthbabyremindbatchsendRequest() *AlibabaalihealthbabyremindbatchsendAPIRequest {
	return &AlibabaalihealthbabyremindbatchsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbabyremindbatchsendAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.baby.remind.batch.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbabyremindbatchsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbabyremindbatchsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchRemindRequests is BatchRemindRequests Setter
// 批量发送提醒
func (r *AlibabaalihealthbabyremindbatchsendAPIRequest) SetBatchRemindRequests(_batchRemindRequests []BatchRemindRequestDto) error {
	r._batchRemindRequests = _batchRemindRequests
	r.Set("batch_remind_requests", _batchRemindRequests)
	return nil
}

// GetBatchRemindRequests BatchRemindRequests Getter
func (r AlibabaalihealthbabyremindbatchsendAPIRequest) GetBatchRemindRequests() []BatchRemindRequestDto {
	return r._batchRemindRequests
}

// SetRemindType is RemindType Setter
// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
func (r *AlibabaalihealthbabyremindbatchsendAPIRequest) SetRemindType(_remindType int64) error {
	r._remindType = _remindType
	r.Set("remind_type", _remindType)
	return nil
}

// GetRemindType RemindType Getter
func (r AlibabaalihealthbabyremindbatchsendAPIRequest) GetRemindType() int64 {
	return r._remindType
}
