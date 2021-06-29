package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
猫超共享库存寄售sopo推送触发 API请求
alibaba.wdk.sopo.push.trigger

猫超共享库存寄售sopo触发推送给商家
*/
type AlibabaWdkSopoPushTriggerRequest struct {
    model.Params
    // 系统自动生成
    wdkOpenPushSoPoRequest   *WdkOpenPushSoPoRequest
}

// 初始化AlibabaWdkSopoPushTriggerRequest对象
func NewAlibabaWdkSopoPushTriggerRequest() *AlibabaWdkSopoPushTriggerRequest{
    return &AlibabaWdkSopoPushTriggerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSopoPushTriggerRequest) GetApiMethodName() string {
    return "alibaba.wdk.sopo.push.trigger"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSopoPushTriggerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WdkOpenPushSoPoRequest Setter
// 系统自动生成
func (r *AlibabaWdkSopoPushTriggerRequest) SetWdkOpenPushSoPoRequest(wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest) error {
    r.wdkOpenPushSoPoRequest = wdkOpenPushSoPoRequest
    r.Set("wdk_open_push_so_po_request", wdkOpenPushSoPoRequest)
    return nil
}

// WdkOpenPushSoPoRequest Getter
func (r AlibabaWdkSopoPushTriggerRequest) GetWdkOpenPushSoPoRequest() *WdkOpenPushSoPoRequest {
    return r.wdkOpenPushSoPoRequest
}
