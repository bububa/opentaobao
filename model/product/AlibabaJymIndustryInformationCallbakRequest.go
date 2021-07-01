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
type AlibabaJymIndustryInformationCallbakAPIRequest struct {
    model.Params
    // 任务ID
    _taskId   string
    // 幂等ID
    _bizId   string
    // 状态
    _status   int64
    // 内容
    _content   string
}

// 初始化AlibabaJymIndustryInformationCallbakAPIRequest对象
func NewAlibabaJymIndustryInformationCallbakRequest() *AlibabaJymIndustryInformationCallbakAPIRequest{
    return &AlibabaJymIndustryInformationCallbakAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetApiMethodName() string {
    return "alibaba.jym.industry.information.callbak"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskId Setter
// 任务ID
func (r *AlibabaJymIndustryInformationCallbakAPIRequest) SetTaskId(_taskId string) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetTaskId() string {
    return r._taskId
}
// BizId Setter
// 幂等ID
func (r *AlibabaJymIndustryInformationCallbakAPIRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetBizId() string {
    return r._bizId
}
// Status Setter
// 状态
func (r *AlibabaJymIndustryInformationCallbakAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetStatus() int64 {
    return r._status
}
// Content Setter
// 内容
func (r *AlibabaJymIndustryInformationCallbakAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaJymIndustryInformationCallbakAPIRequest) GetContent() string {
    return r._content
}
