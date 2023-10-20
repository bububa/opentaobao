package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest 逆向销退入库单入库结果回告 API请求
// alibaba.ascp.uop.supplier.reverseorder.instorage.feedback
//
// ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest struct {
	model.Params
	// 销退单入库结果请求
	_instorageFeedbackRequest *Instoragefeedbackrequest
}

// NewAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest 初始化AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest对象
func NewAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest() *AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest {
	return &AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) Reset() {
	r._instorageFeedbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.instorage.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstorageFeedbackRequest is InstorageFeedbackRequest Setter
// 销退单入库结果请求
func (r *AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) SetInstorageFeedbackRequest(_instorageFeedbackRequest *Instoragefeedbackrequest) error {
	r._instorageFeedbackRequest = _instorageFeedbackRequest
	r.Set("instorage_feedback_request", _instorageFeedbackRequest)
	return nil
}

// GetInstorageFeedbackRequest InstorageFeedbackRequest Getter
func (r AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) GetInstorageFeedbackRequest() *Instoragefeedbackrequest {
	return r._instorageFeedbackRequest
}

var poolAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest()
	},
}

// GetAlibabaAscpUopSupplierReverseorderInstorageFeedbackRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest
func GetAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest() *AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest {
	return poolAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest.Get().(*AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest 将 AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest(v *AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest.Put(v)
}
