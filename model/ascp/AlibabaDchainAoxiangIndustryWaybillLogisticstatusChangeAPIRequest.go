package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) Reset() {
	r._tmsOrderConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeRequest()
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeRequest 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest
func GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest {
	return poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest.Get().(*AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest 将 AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest(v *AlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusChangeAPIRequest.Put(v)
}
