package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest 履约单纬度的仓缺货回告服务 API请求
// alibaba.ascp.uop.supplier.consignorder.outofstock.callback
//
// 商家仓履约单纬度的仓缺货回告接口
type AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest struct {
	model.Params
	// 缺货回告请求模型
	_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest
}

// NewAlibabaascpuopsupplierconsignorderoutofstockcallbackRequest 初始化AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest对象
func NewAlibabaascpuopsupplierconsignorderoutofstockcallbackRequest() *AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest {
	return &AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.outofstock.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderOutofstockCallbackRequest is ConsignorderOutofstockCallbackRequest Setter
// 缺货回告请求模型
func (r *AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest) SetConsignorderOutofstockCallbackRequest(_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest) error {
	r._consignorderOutofstockCallbackRequest = _consignorderOutofstockCallbackRequest
	r.Set("consignorder_outofstock_callback_request", _consignorderOutofstockCallbackRequest)
	return nil
}

// GetConsignorderOutofstockCallbackRequest ConsignorderOutofstockCallbackRequest Getter
func (r AlibabaascpuopsupplierconsignorderoutofstockcallbackAPIRequest) GetConsignorderOutofstockCallbackRequest() *Consignorderoutofstockcallbackrequest {
	return r._consignorderOutofstockCallbackRequest
}
