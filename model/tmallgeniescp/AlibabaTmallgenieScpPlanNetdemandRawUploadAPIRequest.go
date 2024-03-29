package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest 二级物料净需求回传(TL+1) API请求
// alibaba.tmallgenie.scp.plan.netdemand.raw.upload
//
// 二级物料净需求回传(TL+1)
type AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRawRequest *NetDemandRawRequest
}

// NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest 初始化AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest() *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) Reset() {
	r._netDemandRawRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRawRequest is NetDemandRawRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) SetNetDemandRawRequest(_netDemandRawRequest *NetDemandRawRequest) error {
	r._netDemandRawRequest = _netDemandRawRequest
	r.Set("net_demand_raw_request", _netDemandRawRequest)
	return nil
}

// GetNetDemandRawRequest NetDemandRawRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) GetNetDemandRawRequest() *NetDemandRawRequest {
	return r._netDemandRawRequest
}

var poolAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanNetdemandRawUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest
func GetAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest() *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest 将 AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest(v *AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest.Put(v)
}
