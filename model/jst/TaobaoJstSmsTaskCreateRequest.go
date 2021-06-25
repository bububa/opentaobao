package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短信任务创建接口 APIRequest
taobao.jst.sms.task.create

聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
*/
type TaobaoJstSmsTaskCreateRequest struct {
    model.Params

    // 创建任务的入参
    paramCreateSmsTaskRequest   *CreateSmsTaskRequest 

}

func NewTaobaoJstSmsTaskCreateRequest() *TaobaoJstSmsTaskCreateRequest{
    return &TaobaoJstSmsTaskCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsTaskCreateRequest) GetApiMethodName() string {
    return "taobao.jst.sms.task.create"
}

func (r TaobaoJstSmsTaskCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsTaskCreateRequest) SetParamCreateSmsTaskRequest(paramCreateSmsTaskRequest *CreateSmsTaskRequest) error {
    r.paramCreateSmsTaskRequest = paramCreateSmsTaskRequest
    r.Set("param_create_sms_task_request", paramCreateSmsTaskRequest)
    return nil
}

func (r TaobaoJstSmsTaskCreateRequest) GetParamCreateSmsTaskRequest() *CreateSmsTaskRequest {
    return r.paramCreateSmsTaskRequest
}

