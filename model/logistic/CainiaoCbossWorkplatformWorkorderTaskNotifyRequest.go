package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOP-SPI工单任务下发接口 API请求
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV）
*/
type CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest struct {
    model.Params
    // content
    _content   *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct
}

// 初始化CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest对象
func NewCainiaoCbossWorkplatformWorkorderTaskNotifyRequest() *CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest{
    return &CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.task.notify"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// content
func (r *CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest) SetContent(_content *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest) GetContent() *CainiaoCbossWorkplatformWorkorderTaskNotifyStruct {
    return r._content
}
