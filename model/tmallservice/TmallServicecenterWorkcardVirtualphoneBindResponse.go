package tmallservice

import (
    "encoding/xml"

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
    TmallServicecenterWorkcardVirtualphoneBindResponse
}

type TmallServicecenterWorkcardVirtualphoneBindResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_virtualphone_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
