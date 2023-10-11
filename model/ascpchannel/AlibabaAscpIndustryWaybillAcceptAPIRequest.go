package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWaybillAcceptAPIRequest 商家ERP预推单 API请求
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
type AlibabaAscpIndustryWaybillAcceptAPIRequest struct {
	model.Params
	// 开单信息
	_waybillGenRequest *WaybillGenRequest
}

// NewAlibabaAscpIndustryWaybillAcceptRequest 初始化AlibabaAscpIndustryWaybillAcceptAPIRequest对象
func NewAlibabaAscpIndustryWaybillAcceptRequest() *AlibabaAscpIndustryWaybillAcceptAPIRequest {
	return &AlibabaAscpIndustryWaybillAcceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryWaybillAcceptAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.waybill.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryWaybillAcceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryWaybillAcceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillGenRequest is WaybillGenRequest Setter
// 开单信息
func (r *AlibabaAscpIndustryWaybillAcceptAPIRequest) SetWaybillGenRequest(_waybillGenRequest *WaybillGenRequest) error {
	r._waybillGenRequest = _waybillGenRequest
	r.Set("waybill_gen_request", _waybillGenRequest)
	return nil
}

// GetWaybillGenRequest WaybillGenRequest Getter
func (r AlibabaAscpIndustryWaybillAcceptAPIRequest) GetWaybillGenRequest() *WaybillGenRequest {
	return r._waybillGenRequest
}
