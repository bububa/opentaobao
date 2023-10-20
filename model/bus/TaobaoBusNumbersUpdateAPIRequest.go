package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusNumbersUpdateAPIRequest 汽车票车次更新服务 API请求
// taobao.bus.numbers.update
//
// 用于汽车票车次信息的新增、更新和逻辑删除
type TaobaoBusNumbersUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq
}

// NewTaobaoBusNumbersUpdateRequest 初始化TaobaoBusNumbersUpdateAPIRequest对象
func NewTaobaoBusNumbersUpdateRequest() *TaobaoBusNumbersUpdateAPIRequest {
	return &TaobaoBusNumbersUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusNumbersUpdateAPIRequest) Reset() {
	r._paramTopBusNumberUpdateRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusNumbersUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.bus.numbers.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusNumbersUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusNumbersUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopBusNumberUpdateRQ is ParamTopBusNumberUpdateRQ Setter
// 请求参数
func (r *TaobaoBusNumbersUpdateAPIRequest) SetParamTopBusNumberUpdateRQ(_paramTopBusNumberUpdateRQ *TopBusNumberUpdateRq) error {
	r._paramTopBusNumberUpdateRQ = _paramTopBusNumberUpdateRQ
	r.Set("param_top_bus_number_update_r_q", _paramTopBusNumberUpdateRQ)
	return nil
}

// GetParamTopBusNumberUpdateRQ ParamTopBusNumberUpdateRQ Getter
func (r TaobaoBusNumbersUpdateAPIRequest) GetParamTopBusNumberUpdateRQ() *TopBusNumberUpdateRq {
	return r._paramTopBusNumberUpdateRQ
}

var poolTaobaoBusNumbersUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusNumbersUpdateRequest()
	},
}

// GetTaobaoBusNumbersUpdateRequest 从 sync.Pool 获取 TaobaoBusNumbersUpdateAPIRequest
func GetTaobaoBusNumbersUpdateAPIRequest() *TaobaoBusNumbersUpdateAPIRequest {
	return poolTaobaoBusNumbersUpdateAPIRequest.Get().(*TaobaoBusNumbersUpdateAPIRequest)
}

// ReleaseTaobaoBusNumbersUpdateAPIRequest 将 TaobaoBusNumbersUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusNumbersUpdateAPIRequest(v *TaobaoBusNumbersUpdateAPIRequest) {
	v.Reset()
	poolTaobaoBusNumbersUpdateAPIRequest.Put(v)
}
