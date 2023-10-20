package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanMaterielGetAPIRequest 1-IBP同步物料接口 API请求
// alibaba.tmallgenie.scp.plan.materiel.get
//
// IBP同步物料接口
type AlibabaTmallgenieScpPlanMaterielGetAPIRequest struct {
	model.Params
	// 扩展字段
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanMaterielGetRequest 初始化AlibabaTmallgenieScpPlanMaterielGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanMaterielGetRequest() *AlibabaTmallgenieScpPlanMaterielGetAPIRequest {
	return &AlibabaTmallgenieScpPlanMaterielGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanMaterielGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.materiel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展字段
func (r *AlibabaTmallgenieScpPlanMaterielGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanMaterielGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanMaterielGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanMaterielGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanMaterielGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanMaterielGetAPIRequest
func GetAlibabaTmallgenieScpPlanMaterielGetAPIRequest() *AlibabaTmallgenieScpPlanMaterielGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanMaterielGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanMaterielGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanMaterielGetAPIRequest 将 AlibabaTmallgenieScpPlanMaterielGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanMaterielGetAPIRequest(v *AlibabaTmallgenieScpPlanMaterielGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanMaterielGetAPIRequest.Put(v)
}
