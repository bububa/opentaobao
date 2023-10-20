package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) Reset() {
	r._tmsOrderCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.industry.waybill.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangIndustryWaybillCreateRequest()
	},
}

// GetAlibabaDchainAoxiangIndustryWaybillCreateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest
func GetAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest() *AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest {
	return poolAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest.Get().(*AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest 将 AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest(v *AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangIndustryWaybillCreateAPIRequest.Put(v)
}
