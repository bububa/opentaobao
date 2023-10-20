package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksopopushtriggerAPIRequest 猫超共享库存寄售sopo推送触发 API请求
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
type AlibabawdksopopushtriggerAPIRequest struct {
	model.Params
	// 系统自动生成
	_wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest
}

// NewAlibabawdksopopushtriggerRequest 初始化AlibabawdksopopushtriggerAPIRequest对象
func NewAlibabawdksopopushtriggerRequest() *AlibabawdksopopushtriggerAPIRequest {
	return &AlibabawdksopopushtriggerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdksopopushtriggerAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sopo.push.trigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdksopopushtriggerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdksopopushtriggerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWdkOpenPushSoPoRequest is WdkOpenPushSoPoRequest Setter
// 系统自动生成
func (r *AlibabawdksopopushtriggerAPIRequest) SetWdkOpenPushSoPoRequest(_wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest) error {
	r._wdkOpenPushSoPoRequest = _wdkOpenPushSoPoRequest
	r.Set("wdk_open_push_so_po_request", _wdkOpenPushSoPoRequest)
	return nil
}

// GetWdkOpenPushSoPoRequest WdkOpenPushSoPoRequest Getter
func (r AlibabawdksopopushtriggerAPIRequest) GetWdkOpenPushSoPoRequest() *WdkOpenPushSoPoRequest {
	return r._wdkOpenPushSoPoRequest
}
