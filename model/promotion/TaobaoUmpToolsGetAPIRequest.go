package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpToolsGetAPIRequest 查询工具列表 API请求
// taobao.ump.tools.get
//
// 查询工具列表
type TaobaoUmpToolsGetAPIRequest struct {
	model.Params
	// 工具编码
	_toolCode string
}

// NewTaobaoUmpToolsGetRequest 初始化TaobaoUmpToolsGetAPIRequest对象
func NewTaobaoUmpToolsGetRequest() *TaobaoUmpToolsGetAPIRequest {
	return &TaobaoUmpToolsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tools.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ToolCode Setter
// 工具编码
func (r *TaobaoUmpToolsGetAPIRequest) SetToolCode(_toolCode string) error {
	r._toolCode = _toolCode
	r.Set("tool_code", _toolCode)
	return nil
}

// Get ToolCode Getter
func (r TaobaoUmpToolsGetAPIRequest) GetToolCode() string {
	return r._toolCode
}
