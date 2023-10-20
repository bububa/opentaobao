package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest 商家仓wms取消发货反馈回告服务 API请求
// alibaba.ascp.uop.supplier.consignorder.cancel.feedback
//
// 履约单纬度通知商家仓wms取消发货结果反馈回告服务
type AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest struct {
	model.Params
	// 取消发货反馈回告请求
	_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest
}

// NewAlibabaascpuopsupplierconsignordercancelfeedbackRequest 初始化AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest对象
func NewAlibabaascpuopsupplierconsignordercancelfeedbackRequest() *AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest {
	return &AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.cancel.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderCancelFeedbackRequest is ConsignorderCancelFeedbackRequest Setter
// 取消发货反馈回告请求
func (r *AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest) SetConsignorderCancelFeedbackRequest(_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest) error {
	r._consignorderCancelFeedbackRequest = _consignorderCancelFeedbackRequest
	r.Set("consignorder_cancel_feedback_request", _consignorderCancelFeedbackRequest)
	return nil
}

// GetConsignorderCancelFeedbackRequest ConsignorderCancelFeedbackRequest Getter
func (r AlibabaascpuopsupplierconsignordercancelfeedbackAPIRequest) GetConsignorderCancelFeedbackRequest() *Consignordercancelfeedbackrequest {
	return r._consignorderCancelFeedbackRequest
}
