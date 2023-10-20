package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest 12-同步销售数据 API请求
// alibaba.tmallgenie.scp.plan.sale.qty.get
//
// 同步销售数据
type AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanSaleQtyGetRequest 初始化AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanSaleQtyGetRequest() *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest {
	return &AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanSaleQtyGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanSaleQtyGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest
func GetAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest() *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest 将 AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest(v *AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleQtyGetAPIRequest.Put(v)
}
