package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest 客服增加备注 API请求
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
type AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest struct {
	model.Params
	// 物流节点回传入参
	_tmsOrderRemarkRequest *TmsOrderRemarkRequest
}

// NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddRequest 初始化AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest对象
func NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddRequest() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest {
	return &AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) Reset() {
	r._tmsOrderRemarkRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsOrderRemarkRequest is TmsOrderRemarkRequest Setter
// 物流节点回传入参
func (r *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) SetTmsOrderRemarkRequest(_tmsOrderRemarkRequest *TmsOrderRemarkRequest) error {
	r._tmsOrderRemarkRequest = _tmsOrderRemarkRequest
	r.Set("tms_order_remark_request", _tmsOrderRemarkRequest)
	return nil
}

// GetTmsOrderRemarkRequest TmsOrderRemarkRequest Getter
func (r AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) GetTmsOrderRemarkRequest() *TmsOrderRemarkRequest {
	return r._tmsOrderRemarkRequest
}

var poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddRequest()
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddRequest 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest
func GetAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest() *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest {
	return poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest.Get().(*AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest 将 AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest(v *AlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillLogisticstatusRemarkAddAPIRequest.Put(v)
}
