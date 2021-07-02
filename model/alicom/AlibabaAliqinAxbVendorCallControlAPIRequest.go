package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorCallControlAPIRequest 转呼控制接口 API请求
// alibaba.aliqin.axb.vendor.call.control
//
// 转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
type AlibabaAliqinAxbVendorCallControlAPIRequest struct {
	model.Params
	// 转接控制接口request对象
	_startCallRequest *StartCallRequest
}

// NewAlibabaAliqinAxbVendorCallControlRequest 初始化AlibabaAliqinAxbVendorCallControlAPIRequest对象
func NewAlibabaAliqinAxbVendorCallControlRequest() *AlibabaAliqinAxbVendorCallControlAPIRequest {
	return &AlibabaAliqinAxbVendorCallControlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorCallControlAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.call.control"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorCallControlAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartCallRequest is StartCallRequest Setter
// 转接控制接口request对象
func (r *AlibabaAliqinAxbVendorCallControlAPIRequest) SetStartCallRequest(_startCallRequest *StartCallRequest) error {
	r._startCallRequest = _startCallRequest
	r.Set("start_call_request", _startCallRequest)
	return nil
}

// GetStartCallRequest StartCallRequest Getter
func (r AlibabaAliqinAxbVendorCallControlAPIRequest) GetStartCallRequest() *StartCallRequest {
	return r._startCallRequest
}
