package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest 商家仓wms取消发货反馈回告服务 API请求
// alibaba.ascp.uop.supplier.consignorder.cancel.feedback
//
// 履约单纬度通知商家仓wms取消发货结果反馈回告服务
type AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest struct {
	model.Params
	// 取消发货反馈回告请求
	_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest
}

// NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest 初始化AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest() *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) Reset() {
	r._consignorderCancelFeedbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.cancel.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderCancelFeedbackRequest is ConsignorderCancelFeedbackRequest Setter
// 取消发货反馈回告请求
func (r *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) SetConsignorderCancelFeedbackRequest(_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest) error {
	r._consignorderCancelFeedbackRequest = _consignorderCancelFeedbackRequest
	r.Set("consignorder_cancel_feedback_request", _consignorderCancelFeedbackRequest)
	return nil
}

// GetConsignorderCancelFeedbackRequest ConsignorderCancelFeedbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) GetConsignorderCancelFeedbackRequest() *Consignordercancelfeedbackrequest {
	return r._consignorderCancelFeedbackRequest
}

var poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest()
	},
}

// GetAlibabaAscpUopSupplierConsignorderCancelFeedbackRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest
func GetAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest() *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest {
	return poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest.Get().(*AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest 将 AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest(v *AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest.Put(v)
}
