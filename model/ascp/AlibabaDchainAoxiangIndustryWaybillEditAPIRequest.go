package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) Reset() {
	r._tmsOrderUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangIndustryWaybillEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangIndustryWaybillEditRequest()
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillEditRequest 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillEditAPIRequest
func GetAlibabaDchainAoxiangIndustryWaybillEditAPIRequest() *AlibabaDchainAoxiangIndustryWaybillEditAPIRequest {
	return poolAlibabaDchainAoxiangIndustryWaybillEditAPIRequest.Get().(*AlibabaDchainAoxiangIndustryWaybillEditAPIRequest)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillEditAPIRequest 将 AlibabaDchainAoxiangIndustryWaybillEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillEditAPIRequest(v *AlibabaDchainAoxiangIndustryWaybillEditAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillEditAPIRequest.Put(v)
}
