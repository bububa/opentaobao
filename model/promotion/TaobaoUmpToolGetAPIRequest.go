package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpToolGetAPIRequest 查询工具 API请求
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
type TaobaoUmpToolGetAPIRequest struct {
	model.Params
	// 工具的id
	_toolId int64
}

// NewTaobaoUmpToolGetRequest 初始化TaobaoUmpToolGetAPIRequest对象
func NewTaobaoUmpToolGetRequest() *TaobaoUmpToolGetAPIRequest {
	return &TaobaoUmpToolGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tool.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ToolId Setter
// 工具的id
func (r *TaobaoUmpToolGetAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// Get ToolId Getter
func (r TaobaoUmpToolGetAPIRequest) GetToolId() int64 {
	return r._toolId
}
