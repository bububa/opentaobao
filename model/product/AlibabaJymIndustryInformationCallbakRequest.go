package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
VMOS回调行业信息系统 API请求
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统
*/
type AlibabaJymIndustryInformationCallbakRequest struct {
    model.Params
    // 任务ID
    taskId   string
    // 幂等ID
    bizId   string
    // 状态
    status   int64
    // 内容
    content   string
}

// 初始化AlibabaJymIndustryInformationCallbakRequest对象
func NewAlibabaJymIndustryInformationCallbakRequest() *AlibabaJymIndustryInformationCallbakRequest{
    return &AlibabaJymIndustryInformationCallbakRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryInformationCallbakRequest) GetApiMethodName() string {
    return "alibaba.jym.industry.information.callbak"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryInformationCallbakRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskId Setter
// 任务ID
func (r *AlibabaJymIndustryInformationCallbakRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetTaskId() string {
    return r.taskId
}
// BizId Setter
// 幂等ID
func (r *AlibabaJymIndustryInformationCallbakRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

// BizId Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetBizId() string {
    return r.bizId
}
// Status Setter
// 状态
func (r *AlibabaJymIndustryInformationCallbakRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetStatus() int64 {
    return r.status
}
// Content Setter
// 内容
func (r *AlibabaJymIndustryInformationCallbakRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetContent() string {
    return r.content
}
