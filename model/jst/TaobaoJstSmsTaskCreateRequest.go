package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短信任务创建接口 API请求
taobao.jst.sms.task.create

聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
*/
type TaobaoJstSmsTaskCreateRequest struct {
    model.Params
    // 创建任务的入参
    _paramCreateSmsTaskRequest   *CreateSmsTaskRequest
}

// 初始化TaobaoJstSmsTaskCreateRequest对象
func NewTaobaoJstSmsTaskCreateRequest() *TaobaoJstSmsTaskCreateRequest{
    return &TaobaoJstSmsTaskCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTaskCreateRequest) GetApiMethodName() string {
    return "taobao.jst.sms.task.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCreateSmsTaskRequest Setter
// 创建任务的入参
func (r *TaobaoJstSmsTaskCreateRequest) SetParamCreateSmsTaskRequest(_paramCreateSmsTaskRequest *CreateSmsTaskRequest) error {
    r._paramCreateSmsTaskRequest = _paramCreateSmsTaskRequest
    r.Set("param_create_sms_task_request", _paramCreateSmsTaskRequest)
    return nil
}

// ParamCreateSmsTaskRequest Getter
func (r TaobaoJstSmsTaskCreateRequest) GetParamCreateSmsTaskRequest() *CreateSmsTaskRequest {
    return r._paramCreateSmsTaskRequest
}
