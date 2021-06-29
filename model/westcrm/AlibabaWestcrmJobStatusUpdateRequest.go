package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新工单状态 API请求
alibaba.westcrm.job.status.update

更新工单处理状态
*/
type AlibabaWestcrmJobStatusUpdateRequest struct {
    model.Params
    // 状态
    _status   int64
    // 园区id
    _campusId   int64
    // 反馈id
    _crmComplaintId   int64
    // 回复内容
    _replyContent   string
}

// 初始化AlibabaWestcrmJobStatusUpdateRequest对象
func NewAlibabaWestcrmJobStatusUpdateRequest() *AlibabaWestcrmJobStatusUpdateRequest{
    return &AlibabaWestcrmJobStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmJobStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.westcrm.job.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmJobStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 状态
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetStatus() int64 {
    return r._status
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetCampusId() int64 {
    return r._campusId
}
// CrmComplaintId Setter
// 反馈id
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetCrmComplaintId(_crmComplaintId int64) error {
    r._crmComplaintId = _crmComplaintId
    r.Set("crm_complaint_id", _crmComplaintId)
    return nil
}

// CrmComplaintId Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetCrmComplaintId() int64 {
    return r._crmComplaintId
}
// ReplyContent Setter
// 回复内容
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetReplyContent(_replyContent string) error {
    r._replyContent = _replyContent
    r.Set("reply_content", _replyContent)
    return nil
}

// ReplyContent Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetReplyContent() string {
    return r._replyContent
}
