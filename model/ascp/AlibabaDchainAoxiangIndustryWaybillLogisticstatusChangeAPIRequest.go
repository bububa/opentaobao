package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest 物流节点回传 API请求
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest struct {
	model.Params
	// 物流节点回传入参
	_tmsOrderConfirmRequest *TmsOrderConfirmRequest
}

// NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeRequest 初始化AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest对象
func NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeRequest() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest {
	return &AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTmsOrderConfirmRequest is TmsOrderConfirmRequest Setter
// 物流节点回传入参
func (r *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) SetTmsOrderConfirmRequest(_tmsOrderConfirmRequest *TmsOrderConfirmRequest) error {
	r._tmsOrderConfirmRequest = _tmsOrderConfirmRequest
	r.Set("tms_order_confirm_request", _tmsOrderConfirmRequest)
	return nil
}

// GetTmsOrderConfirmRequest TmsOrderConfirmRequest Getter
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetTmsOrderConfirmRequest() *TmsOrderConfirmRequest {
	return r._tmsOrderConfirmRequest
}
