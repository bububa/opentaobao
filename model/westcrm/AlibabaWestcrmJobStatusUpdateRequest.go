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
    status   int64
    // 园区id
    campusId   int64
    // 反馈id
    crmComplaintId   int64
    // 回复内容
    replyContent   string
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
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetStatus() int64 {
    return r.status
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetCampusId() int64 {
    return r.campusId
}
// CrmComplaintId Setter
// 反馈id
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetCrmComplaintId(crmComplaintId int64) error {
    r.crmComplaintId = crmComplaintId
    r.Set("crm_complaint_id", crmComplaintId)
    return nil
}

// CrmComplaintId Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetCrmComplaintId() int64 {
    return r.crmComplaintId
}
// ReplyContent Setter
// 回复内容
func (r *AlibabaWestcrmJobStatusUpdateRequest) SetReplyContent(replyContent string) error {
    r.replyContent = replyContent
    r.Set("reply_content", replyContent)
    return nil
}

// ReplyContent Getter
func (r AlibabaWestcrmJobStatusUpdateRequest) GetReplyContent() string {
    return r.replyContent
}
