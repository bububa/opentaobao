package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrywaybillpreacceptAPIRequest 商家ERP预推单 API请求
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
type AlibabaascpindustrywaybillpreacceptAPIRequest struct {
	model.Params
	// 开单信息
	_waybillGenRequest *WaybillGenRequest
}

// NewAlibabaascpindustrywaybillpreacceptRequest 初始化AlibabaascpindustrywaybillpreacceptAPIRequest对象
func NewAlibabaascpindustrywaybillpreacceptRequest() *AlibabaascpindustrywaybillpreacceptAPIRequest {
	return &AlibabaascpindustrywaybillpreacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustrywaybillpreacceptAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.waybill.pre.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustrywaybillpreacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustrywaybillpreacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillGenRequest is WaybillGenRequest Setter
// 开单信息
func (r *AlibabaascpindustrywaybillpreacceptAPIRequest) SetWaybillGenRequest(_waybillGenRequest *WaybillGenRequest) error {
	r._waybillGenRequest = _waybillGenRequest
	r.Set("waybill_gen_request", _waybillGenRequest)
	return nil
}

// GetWaybillGenRequest WaybillGenRequest Getter
func (r AlibabaascpindustrywaybillpreacceptAPIRequest) GetWaybillGenRequest() *WaybillGenRequest {
	return r._waybillGenRequest
}
