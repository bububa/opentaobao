package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusBusnumberGetAPIRequest 汽车票车次查询 API请求
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
type TaobaoBusBusnumberGetAPIRequest struct {
	model.Params
	// 车次查询入参
	_paramBusNumberSearchRQ *BusNumberSearchRq
}

// NewTaobaoBusBusnumberGetRequest 初始化TaobaoBusBusnumberGetAPIRequest对象
func NewTaobaoBusBusnumberGetRequest() *TaobaoBusBusnumberGetAPIRequest {
	return &TaobaoBusBusnumberGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusBusnumberGetAPIRequest) Reset() {
	r._paramBusNumberSearchRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusBusnumberGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.busnumber.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusBusnumberGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusBusnumberGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBusNumberSearchRQ is ParamBusNumberSearchRQ Setter
// 车次查询入参
func (r *TaobaoBusBusnumberGetAPIRequest) SetParamBusNumberSearchRQ(_paramBusNumberSearchRQ *BusNumberSearchRq) error {
	r._paramBusNumberSearchRQ = _paramBusNumberSearchRQ
	r.Set("param_bus_number_search_r_q", _paramBusNumberSearchRQ)
	return nil
}

// GetParamBusNumberSearchRQ ParamBusNumberSearchRQ Getter
func (r TaobaoBusBusnumberGetAPIRequest) GetParamBusNumberSearchRQ() *BusNumberSearchRq {
	return r._paramBusNumberSearchRQ
}

var poolTaobaoBusBusnumberGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusBusnumberGetRequest()
	},
}

// GetTaobaoBusBusnumberGetRequest 从 sync.Pool 获取 TaobaoBusBusnumberGetAPIRequest
func GetTaobaoBusBusnumberGetAPIRequest() *TaobaoBusBusnumberGetAPIRequest {
	return poolTaobaoBusBusnumberGetAPIRequest.Get().(*TaobaoBusBusnumberGetAPIRequest)
}

// ReleaseTaobaoBusBusnumberGetAPIRequest 将 TaobaoBusBusnumberGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusBusnumberGetAPIRequest(v *TaobaoBusBusnumberGetAPIRequest) {
	v.Reset()
	poolTaobaoBusBusnumberGetAPIRequest.Put(v)
}
