package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOP-SPI工单任务下发接口 APIRequest
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV）
*/
type CainiaoCbossWorkplatformWorkorderTaskNotifyRequest struct {
    model.Params

    // content
    content   *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct 

}

func NewCainiaoCbossWorkplatformWorkorderTaskNotifyRequest() *CainiaoCbossWorkplatformWorkorderTaskNotifyRequest{
    return &CainiaoCbossWorkplatformWorkorderTaskNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformWorkorderTaskNotifyRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.task.notify"
}

func (r CainiaoCbossWorkplatformWorkorderTaskNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformWorkorderTaskNotifyRequest) SetContent(content *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderTaskNotifyRequest) GetContent() *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct {
    return r.content
}

