package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest 逆向销退入库单到仓结果回告 API请求
// alibaba.ascp.uop.supplier.reverseorder.instorage.result
//
// ERP回告销退入库单到仓信息回告
type AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest struct {
	model.Params
	// 消退入库单结果请求
	_instorageResultRequest *Instorageresultrequest
}

// NewAlibabaascpuopsupplierreverseorderinstorageresultRequest 初始化AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest对象
func NewAlibabaascpuopsupplierreverseorderinstorageresultRequest() *AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest {
	return &AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.reverseorder.instorage.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstorageResultRequest is InstorageResultRequest Setter
// 消退入库单结果请求
func (r *AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest) SetInstorageResultRequest(_instorageResultRequest *Instorageresultrequest) error {
	r._instorageResultRequest = _instorageResultRequest
	r.Set("instorage_result_request", _instorageResultRequest)
	return nil
}

// GetInstorageResultRequest InstorageResultRequest Getter
func (r AlibabaascpuopsupplierreverseorderinstorageresultAPIRequest) GetInstorageResultRequest() *Instorageresultrequest {
	return r._instorageResultRequest
}
