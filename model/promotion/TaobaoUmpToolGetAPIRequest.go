package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumptoolgetAPIRequest 查询工具 API请求
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
type TaobaoumptoolgetAPIRequest struct {
	model.Params
	// 工具的id
	_toolId int64
}

// NewTaobaoumptoolgetRequest 初始化TaobaoumptoolgetAPIRequest对象
func NewTaobaoumptoolgetRequest() *TaobaoumptoolgetAPIRequest {
	return &TaobaoumptoolgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumptoolgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tool.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumptoolgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumptoolgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToolId is ToolId Setter
// 工具的id
func (r *TaobaoumptoolgetAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoumptoolgetAPIRequest) GetToolId() int64 {
	return r._toolId
}
