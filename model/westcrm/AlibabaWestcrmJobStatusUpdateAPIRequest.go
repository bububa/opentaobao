package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmJobStatusUpdateAPIRequest 更新工单状态 API请求
// alibaba.westcrm.job.status.update
//
// 更新工单处理状态
type AlibabaWestcrmJobStatusUpdateAPIRequest struct {
	model.Params
	// 回复内容
	_replyContent string
	// 状态
	_status int64
	// 园区id
	_campusId int64
	// 反馈id
	_crmComplaintId int64
}

// NewAlibabaWestcrmJobStatusUpdateRequest 初始化AlibabaWestcrmJobStatusUpdateAPIRequest对象
func NewAlibabaWestcrmJobStatusUpdateRequest() *AlibabaWestcrmJobStatusUpdateAPIRequest {
	return &AlibabaWestcrmJobStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.job.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReplyContent is ReplyContent Setter
// 回复内容
func (r *AlibabaWestcrmJobStatusUpdateAPIRequest) SetReplyContent(_replyContent string) error {
	r._replyContent = _replyContent
	r.Set("reply_content", _replyContent)
	return nil
}

// GetReplyContent ReplyContent Getter
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetReplyContent() string {
	return r._replyContent
}

// SetStatus is Status Setter
// 状态
func (r *AlibabaWestcrmJobStatusUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaWestcrmJobStatusUpdateAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetCrmComplaintId is CrmComplaintId Setter
// 反馈id
func (r *AlibabaWestcrmJobStatusUpdateAPIRequest) SetCrmComplaintId(_crmComplaintId int64) error {
	r._crmComplaintId = _crmComplaintId
	r.Set("crm_complaint_id", _crmComplaintId)
	return nil
}

// GetCrmComplaintId CrmComplaintId Getter
func (r AlibabaWestcrmJobStatusUpdateAPIRequest) GetCrmComplaintId() int64 {
	return r._crmComplaintId
}
