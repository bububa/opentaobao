package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单系统的工单进度下发 API请求
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI）
*/
type CainiaoCbossWorkplatformWorkorderProcessNotifyRequest struct {
    model.Params
    // 服务入参
    content   *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct
}

// 初始化CainiaoCbossWorkplatformWorkorderProcessNotifyRequest对象
func NewCainiaoCbossWorkplatformWorkorderProcessNotifyRequest() *CainiaoCbossWorkplatformWorkorderProcessNotifyRequest{
    return &CainiaoCbossWorkplatformWorkorderProcessNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.process.notify"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 服务入参
func (r *CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) SetContent(content *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyRequest) GetContent() *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct {
    return r.content
}
