package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkspreadgetAPIRequest 淘宝客-公用-长链转短链 API请求
// taobao.tbk.spread.get
//
// 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
// 现阶段只支持短连接。
type TaobaotbkspreadgetAPIRequest struct {
	model.Params
	// 请求列表，内部包含多个url
	_requests []TbkSpreadRequest
}

// NewTaobaotbkspreadgetRequest 初始化TaobaotbkspreadgetAPIRequest对象
func NewTaobaotbkspreadgetRequest() *TaobaotbkspreadgetAPIRequest {
	return &TaobaotbkspreadgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkspreadgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.spread.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkspreadgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkspreadgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequests is Requests Setter
// 请求列表，内部包含多个url
func (r *TaobaotbkspreadgetAPIRequest) SetRequests(_requests []TbkSpreadRequest) error {
	r._requests = _requests
	r.Set("requests", _requests)
	return nil
}

// GetRequests Requests Getter
func (r TaobaotbkspreadgetAPIRequest) GetRequests() []TbkSpreadRequest {
	return r._requests
}
