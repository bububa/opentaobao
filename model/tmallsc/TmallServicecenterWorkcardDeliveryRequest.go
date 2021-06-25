package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开始配送工单 APIRequest
tmall.servicecenter.workcard.delivery

服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
*/
type TmallServicecenterWorkcardDeliveryRequest struct {
    model.Params

    // 工单配送请求参数
    identifyTaskDeliveryRequest   *IdentifyTaskDeliveryRequest 

}

func NewTmallServicecenterWorkcardDeliveryRequest() *TmallServicecenterWorkcardDeliveryRequest{
    return &TmallServicecenterWorkcardDeliveryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardDeliveryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.delivery"
}

func (r TmallServicecenterWorkcardDeliveryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardDeliveryRequest) SetIdentifyTaskDeliveryRequest(identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest) error {
    r.identifyTaskDeliveryRequest = identifyTaskDeliveryRequest
    r.Set("identify_task_delivery_request", identifyTaskDeliveryRequest)
    return nil
}

func (r TmallServicecenterWorkcardDeliveryRequest) GetIdentifyTaskDeliveryRequest() *IdentifyTaskDeliveryRequest {
    return r.identifyTaskDeliveryRequest
}

