package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest 服务商开运单 API请求
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
type AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest struct {
	model.Params
	// 服务商开单
	_tmsOrderCreateRequest *TmsOrderCreateRequest
}

// NewAlibabaDchainAoxiangIndustryWaybillCreateRequest 初始化AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest对象
func NewAlibabaDchainAoxiangIndustryWaybillCreateRequest() *AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest {
	return &AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTmsOrderCreateRequest is TmsOrderCreateRequest Setter
// 服务商开单
func (r *AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) SetTmsOrderCreateRequest(_tmsOrderCreateRequest *TmsOrderCreateRequest) error {
	r._tmsOrderCreateRequest = _tmsOrderCreateRequest
	r.Set("tms_order_create_request", _tmsOrderCreateRequest)
	return nil
}

// GetTmsOrderCreateRequest TmsOrderCreateRequest Getter
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetTmsOrderCreateRequest() *TmsOrderCreateRequest {
	return r._tmsOrderCreateRequest
}
