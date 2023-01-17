package bus

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
