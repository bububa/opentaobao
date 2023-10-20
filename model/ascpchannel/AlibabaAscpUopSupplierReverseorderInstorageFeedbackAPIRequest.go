package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest 逆向销退入库单入库结果回告 API请求
// alibaba.ascp.uop.supplier.reverseorder.instorage.feedback
//
// ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
type AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest struct {
	model.Params
	// 销退单入库结果请求
	_instorageFeedbackRequest *Instoragefeedbackrequest
}

// NewAlibabaascpuopsupplierreverseorderinstoragefeedbackRequest 初始化AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest对象
func NewAlibabaascpuopsupplierreverseorderinstoragefeedbackRequest() *AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest {
	return &AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.instorage.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstorageFeedbackRequest is InstorageFeedbackRequest Setter
// 销退单入库结果请求
func (r *AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest) SetInstorageFeedbackRequest(_instorageFeedbackRequest *Instoragefeedbackrequest) error {
	r._instorageFeedbackRequest = _instorageFeedbackRequest
	r.Set("instorage_feedback_request", _instorageFeedbackRequest)
	return nil
}

// GetInstorageFeedbackRequest InstorageFeedbackRequest Getter
func (r AlibabaascpuopsupplierreverseorderinstoragefeedbackAPIRequest) GetInstorageFeedbackRequest() *Instoragefeedbackrequest {
	return r._instorageFeedbackRequest
}
