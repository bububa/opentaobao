package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillEditAPIRequest 服务商编辑运单 API请求
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
type AlibabaDchainAoxiangIndustryWaybillEditAPIRequest struct {
	model.Params
	// 服务商编辑
	_tmsOrderUpdateRequest *TmsOrderUpdateRequest
}

// NewAlibabaDchainAoxiangIndustryWaybillEditRequest 初始化AlibabaDchainAoxiangIndustryWaybillEditAPIRequest对象
func NewAlibabaDchainAoxiangIndustryWaybillEditRequest() *AlibabaDchainAoxiangIndustryWaybillEditAPIRequest {
	return &AlibabaDchainAoxiangIndustryWaybillEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTmsOrderUpdateRequest is TmsOrderUpdateRequest Setter
// 服务商编辑
func (r *AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) SetTmsOrderUpdateRequest(_tmsOrderUpdateRequest *TmsOrderUpdateRequest) error {
	r._tmsOrderUpdateRequest = _tmsOrderUpdateRequest
	r.Set("tms_order_update_request", _tmsOrderUpdateRequest)
	return nil
}

// GetTmsOrderUpdateRequest TmsOrderUpdateRequest Getter
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetTmsOrderUpdateRequest() *TmsOrderUpdateRequest {
	return r._tmsOrderUpdateRequest
}
