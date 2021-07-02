package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianServerDateGetAPIRequest 服务器时间获取 API请求
// taobao.baodian.server.date.get
//
// 获取服务器时间
type TaobaoBaodianServerDateGetAPIRequest struct {
	model.Params
}

// NewTaobaoBaodianServerDateGetRequest 初始化TaobaoBaodianServerDateGetAPIRequest对象
func NewTaobaoBaodianServerDateGetRequest() *TaobaoBaodianServerDateGetAPIRequest {
	return &TaobaoBaodianServerDateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaodianServerDateGetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.server.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaodianServerDateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
