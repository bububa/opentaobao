package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单维度虚拟中间号绑定 API请求
tmall.servicecenter.workcard.virtualphone.bind

服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
叫消费者的手机号虚拟号码给到客服。
*/
type TmallServicecenterWorkcardVirtualphoneBindRequest struct {
    model.Params
    // 绑定阿里通讯号入参
    workcardRequest   *WorkcardBaseRequest
}

// 初始化TmallServicecenterWorkcardVirtualphoneBindRequest对象
func NewTmallServicecenterWorkcardVirtualphoneBindRequest() *TmallServicecenterWorkcardVirtualphoneBindRequest{
    return &TmallServicecenterWorkcardVirtualphoneBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.virtualphone.bind"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardRequest Setter
// 绑定阿里通讯号入参
func (r *TmallServicecenterWorkcardVirtualphoneBindRequest) SetWorkcardRequest(workcardRequest *WorkcardBaseRequest) error {
    r.workcardRequest = workcardRequest
    r.Set("workcard_request", workcardRequest)
    return nil
}

// WorkcardRequest Getter
func (r TmallServicecenterWorkcardVirtualphoneBindRequest) GetWorkcardRequest() *WorkcardBaseRequest {
    return r.workcardRequest
}
