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
    _taskId   string
    // 幂等ID
    _bizId   string
    // 状态
    _status   int64
    // 内容
    _content   string
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
func (r *AlibabaJymIndustryInformationCallbakRequest) SetTaskId(_taskId string) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetTaskId() string {
    return r._taskId
}
// BizId Setter
// 幂等ID
func (r *AlibabaJymIndustryInformationCallbakRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetBizId() string {
    return r._bizId
}
// Status Setter
// 状态
func (r *AlibabaJymIndustryInformationCallbakRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetStatus() int64 {
    return r._status
}
// Content Setter
// 内容
func (r *AlibabaJymIndustryInformationCallbakRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaJymIndustryInformationCallbakRequest) GetContent() string {
    return r._content
}
