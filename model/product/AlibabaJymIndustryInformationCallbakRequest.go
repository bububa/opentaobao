package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
VMOS回调行业信息系统 APIRequest
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

func NewAlibabaJymIndustryInformationCallbakRequest() *AlibabaJymIndustryInformationCallbakRequest{
    return &AlibabaJymIndustryInformationCallbakRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetApiMethodName() string {
    return "alibaba.jym.industry.information.callbak"
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaJymIndustryInformationCallbakRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetTaskId() string {
    return r.taskId
}

func (r *AlibabaJymIndustryInformationCallbakRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetBizId() string {
    return r.bizId
}

func (r *AlibabaJymIndustryInformationCallbakRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaJymIndustryInformationCallbakRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaJymIndustryInformationCallbakRequest) GetContent() string {
    return r.content
}

