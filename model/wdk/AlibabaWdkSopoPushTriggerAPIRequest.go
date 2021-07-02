package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSopoPushTriggerAPIRequest 猫超共享库存寄售sopo推送触发 API请求
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
type AlibabaWdkSopoPushTriggerAPIRequest struct {
	model.Params
	// 系统自动生成
	_wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest
}

// NewAlibabaWdkSopoPushTriggerRequest 初始化AlibabaWdkSopoPushTriggerAPIRequest对象
func NewAlibabaWdkSopoPushTriggerRequest() *AlibabaWdkSopoPushTriggerAPIRequest {
	return &AlibabaWdkSopoPushTriggerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSopoPushTriggerAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sopo.push.trigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSopoPushTriggerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWdkOpenPushSoPoRequest is WdkOpenPushSoPoRequest Setter
// 系统自动生成
func (r *AlibabaWdkSopoPushTriggerAPIRequest) SetWdkOpenPushSoPoRequest(_wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest) error {
	r._wdkOpenPushSoPoRequest = _wdkOpenPushSoPoRequest
	r.Set("wdk_open_push_so_po_request", _wdkOpenPushSoPoRequest)
	return nil
}

// GetWdkOpenPushSoPoRequest WdkOpenPushSoPoRequest Getter
func (r AlibabaWdkSopoPushTriggerAPIRequest) GetWdkOpenPushSoPoRequest() *WdkOpenPushSoPoRequest {
	return r._wdkOpenPushSoPoRequest
}
