package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest 10-同步库存现有量 API请求
// alibaba.tmallgenie.scp.plan.inventor.qty.get
//
// 同步库存现有量
type AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanInventorQtyGetRequest 初始化AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanInventorQtyGetRequest() *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest {
	return &AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.inventor.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanInventorQtyGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanInventorQtyGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest
func GetAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest() *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest 将 AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest(v *AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanInventorQtyGetAPIRequest.Put(v)
}
