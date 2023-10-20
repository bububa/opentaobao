package alihealthcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBabyRemindBatchSendAPIRequest 批量发送母婴提醒 API请求
// alibaba.alihealth.baby.remind.batch.send
//
// 批量发送母婴提醒
type AlibabaAlihealthBabyRemindBatchSendAPIRequest struct {
	model.Params
	// 批量发送提醒
	_batchRemindRequests []BatchRemindRequestDto
	// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
	_remindType int64
}

// NewAlibabaAlihealthBabyRemindBatchSendRequest 初始化AlibabaAlihealthBabyRemindBatchSendAPIRequest对象
func NewAlibabaAlihealthBabyRemindBatchSendRequest() *AlibabaAlihealthBabyRemindBatchSendAPIRequest {
	return &AlibabaAlihealthBabyRemindBatchSendAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBabyRemindBatchSendAPIRequest) Reset() {
	r._batchRemindRequests = r._batchRemindRequests[:0]
	r._remindType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.baby.remind.batch.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchRemindRequests is BatchRemindRequests Setter
// 批量发送提醒
func (r *AlibabaAlihealthBabyRemindBatchSendAPIRequest) SetBatchRemindRequests(_batchRemindRequests []BatchRemindRequestDto) error {
	r._batchRemindRequests = _batchRemindRequests
	r.Set("batch_remind_requests", _batchRemindRequests)
	return nil
}

// GetBatchRemindRequests BatchRemindRequests Getter
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetBatchRemindRequests() []BatchRemindRequestDto {
	return r._batchRemindRequests
}

// SetRemindType is RemindType Setter
// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
func (r *AlibabaAlihealthBabyRemindBatchSendAPIRequest) SetRemindType(_remindType int64) error {
	r._remindType = _remindType
	r.Set("remind_type", _remindType)
	return nil
}

// GetRemindType RemindType Getter
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetRemindType() int64 {
	return r._remindType
}

var poolAlibabaAlihealthBabyRemindBatchSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBabyRemindBatchSendRequest()
	},
}

// GetAlibabaAlihealthBabyRemindBatchSendRequest 从 sync.Pool 获取 AlibabaAlihealthBabyRemindBatchSendAPIRequest
func GetAlibabaAlihealthBabyRemindBatchSendAPIRequest() *AlibabaAlihealthBabyRemindBatchSendAPIRequest {
	return poolAlibabaAlihealthBabyRemindBatchSendAPIRequest.Get().(*AlibabaAlihealthBabyRemindBatchSendAPIRequest)
}

// ReleaseAlibabaAlihealthBabyRemindBatchSendAPIRequest 将 AlibabaAlihealthBabyRemindBatchSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBabyRemindBatchSendAPIRequest(v *AlibabaAlihealthBabyRemindBatchSendAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBabyRemindBatchSendAPIRequest.Put(v)
}
