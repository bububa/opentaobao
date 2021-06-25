package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单系统的工单进度下发 APIRequest
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI）
*/
type CainiaoCbossWorkplatformWorkorderProcessNotifyRequest struct {
    model.Params

    // 服务入参
    content   *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct 

}

func NewCainiaoCbossWorkplatformWorkorderProcessNotifyRequest() *CainiaoCbossWorkplatformWorkorderProcessNotifyRequest{
    return &CainiaoCbossWorkplatformWorkorderProcessNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.process.notify"
}

func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) SetContent(content *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetContent() *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct {
    return r.content
}

