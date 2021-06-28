package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单维度虚拟中间号绑定 APIResponse
tmall.servicecenter.workcard.virtualphone.bind

服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
叫消费者的手机号虚拟号码给到客服。
*/
type TmallServicecenterWorkcardVirtualphoneBindAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardVirtualphoneBindResponse `json:"tmall_servicecenter_workcard_virtualphone_bind_response,omitempty"` 
    TmallServicecenterWorkcardVirtualphoneBindResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardVirtualphoneBindResponse struct {

    // 系统自动生成
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardVirtualphoneBindResponse struct {

    // 系统自动生成
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
