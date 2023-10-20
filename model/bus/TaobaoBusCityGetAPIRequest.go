package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusCityGetAPIRequest 城市接口 API请求
// taobao.bus.city.get
//
// 汽车票出发城市获取接口，获取所有出发城市
type TaobaoBusCityGetAPIRequest struct {
	model.Params
}

// NewTaobaoBusCityGetRequest 初始化TaobaoBusCityGetAPIRequest对象
func NewTaobaoBusCityGetRequest() *TaobaoBusCityGetAPIRequest {
	return &TaobaoBusCityGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusCityGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusCityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusCityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoBusCityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusCityGetRequest()
	},
}

// GetTaobaoBusCityGetRequest 从 sync.Pool 获取 TaobaoBusCityGetAPIRequest
func GetTaobaoBusCityGetAPIRequest() *TaobaoBusCityGetAPIRequest {
	return poolTaobaoBusCityGetAPIRequest.Get().(*TaobaoBusCityGetAPIRequest)
}

// ReleaseTaobaoBusCityGetAPIRequest 将 TaobaoBusCityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusCityGetAPIRequest(v *TaobaoBusCityGetAPIRequest) {
	v.Reset()
	poolTaobaoBusCityGetAPIRequest.Put(v)
}
