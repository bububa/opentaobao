package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumptoolsgetAPIRequest 查询工具列表 API请求
// taobao.ump.tools.get
//
// 查询工具列表
type TaobaoumptoolsgetAPIRequest struct {
	model.Params
	// 工具编码
	_toolCode string
}

// NewTaobaoumptoolsgetRequest 初始化TaobaoumptoolsgetAPIRequest对象
func NewTaobaoumptoolsgetRequest() *TaobaoumptoolsgetAPIRequest {
	return &TaobaoumptoolsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumptoolsgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tools.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumptoolsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumptoolsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToolCode is ToolCode Setter
// 工具编码
func (r *TaobaoumptoolsgetAPIRequest) SetToolCode(_toolCode string) error {
	r._toolCode = _toolCode
	r.Set("tool_code", _toolCode)
	return nil
}

// GetToolCode ToolCode Getter
func (r TaobaoumptoolsgetAPIRequest) GetToolCode() string {
	return r._toolCode
}
