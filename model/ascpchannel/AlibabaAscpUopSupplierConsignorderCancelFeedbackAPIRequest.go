package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest
商家仓wms取消发货反馈回告服务 API请求
alibaba.ascp.uop.supplier.consignorder.cancel.feedback

履约单纬度通知商家仓wms取消发货结果反馈回告服务 */
type AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest struct {
	model.Params
	// 取消发货反馈回告请求
	_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest
}

// NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest 初始化AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest() *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.cancel.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ConsignorderCancelFeedbackRequest Setter
// 取消发货反馈回告请求
func (r *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) SetConsignorderCancelFeedbackRequest(_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest) error {
	r._consignorderCancelFeedbackRequest = _consignorderCancelFeedbackRequest
	r.Set("consignorder_cancel_feedback_request", _consignorderCancelFeedbackRequest)
	return nil
}

// Get ConsignorderCancelFeedbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetConsignorderCancelFeedbackRequest() *Consignordercancelfeedbackrequest {
	return r._consignorderCancelFeedbackRequest
}
