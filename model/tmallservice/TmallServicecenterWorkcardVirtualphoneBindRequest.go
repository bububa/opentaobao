package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单维度虚拟中间号绑定 APIRequest
tmall.servicecenter.workcard.virtualphone.bind

服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
叫消费者的手机号虚拟号码给到客服。
*/
type TmallServicecenterWorkcardVirtualphoneBindRequest struct {
    model.Params

    // 绑定阿里通讯号入参
    workcardRequest   *WorkcardBaseRequest 

}

func NewTmallServicecenterWorkcardVirtualphoneBindRequest() *TmallServicecenterWorkcardVirtualphoneBindRequest{
    return &TmallServicecenterWorkcardVirtualphoneBindRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.virtualphone.bind"
}

func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardVirtualphoneBindRequest) SetWorkcardRequest(workcardRequest *WorkcardBaseRequest) error {
    r.workcardRequest = workcardRequest
    r.Set("workcard_request", workcardRequest)
    return nil
}

func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetWorkcardRequest() *WorkcardBaseRequest {
    return r.workcardRequest
}

