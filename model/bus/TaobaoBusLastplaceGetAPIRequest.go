package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusLastplaceGetAPIRequest 获取目的地数据 API请求
// taobao.bus.lastplace.get
//
// 传入城市 获取对应的目的地
type TaobaoBusLastplaceGetAPIRequest struct {
	model.Params
	// 目的地查询参数
	_paramLastPlaceSearchRQ *ParamLastPlaceSearchRq
}

// NewTaobaoBusLastplaceGetRequest 初始化TaobaoBusLastplaceGetAPIRequest对象
func NewTaobaoBusLastplaceGetRequest() *TaobaoBusLastplaceGetAPIRequest {
	return &TaobaoBusLastplaceGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusLastplaceGetAPIRequest) Reset() {
	r._paramLastPlaceSearchRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusLastplaceGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.lastplace.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusLastplaceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusLastplaceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLastPlaceSearchRQ is ParamLastPlaceSearchRQ Setter
// 目的地查询参数
func (r *TaobaoBusLastplaceGetAPIRequest) SetParamLastPlaceSearchRQ(_paramLastPlaceSearchRQ *ParamLastPlaceSearchRq) error {
	r._paramLastPlaceSearchRQ = _paramLastPlaceSearchRQ
	r.Set("param_last_place_search_r_q", _paramLastPlaceSearchRQ)
	return nil
}

// GetParamLastPlaceSearchRQ ParamLastPlaceSearchRQ Getter
func (r TaobaoBusLastplaceGetAPIRequest) GetParamLastPlaceSearchRQ() *ParamLastPlaceSearchRq {
	return r._paramLastPlaceSearchRQ
}

var poolTaobaoBusLastplaceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusLastplaceGetRequest()
	},
}

// GetTaobaoBusLastplaceGetRequest 从 sync.Pool 获取 TaobaoBusLastplaceGetAPIRequest
func GetTaobaoBusLastplaceGetAPIRequest() *TaobaoBusLastplaceGetAPIRequest {
	return poolTaobaoBusLastplaceGetAPIRequest.Get().(*TaobaoBusLastplaceGetAPIRequest)
}

// ReleaseTaobaoBusLastplaceGetAPIRequest 将 TaobaoBusLastplaceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusLastplaceGetAPIRequest(v *TaobaoBusLastplaceGetAPIRequest) {
	v.Reset()
	poolTaobaoBusLastplaceGetAPIRequest.Put(v)
}
