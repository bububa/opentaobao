package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest 23-Net Demand净需求回传 API请求
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
type AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanNetdemandUploadRequest 初始化AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanNetdemandUploadRequest() *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) Reset() {
	r._netDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.netdemand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}

var poolAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanNetdemandUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanNetdemandUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest
func GetAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest() *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest 将 AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest(v *AlibabaTmallgenieScpPlanNetdemandUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanNetdemandUploadAPIRequest.Put(v)
}
