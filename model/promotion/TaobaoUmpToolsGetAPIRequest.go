package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpToolsGetAPIRequest) Reset() {
	r._toolCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tools.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpToolsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToolCode is ToolCode Setter
// 工具编码
func (r *TaobaoUmpToolsGetAPIRequest) SetToolCode(_toolCode string) error {
	r._toolCode = _toolCode
	r.Set("tool_code", _toolCode)
	return nil
}

// GetToolCode ToolCode Getter
func (r TaobaoUmpToolsGetAPIRequest) GetToolCode() string {
	return r._toolCode
}

var poolTaobaoUmpToolsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpToolsGetRequest()
	},
}

// GetTaobaoUmpToolsGetRequest 从 sync.Pool 获取 TaobaoUmpToolsGetAPIRequest
func GetTaobaoUmpToolsGetAPIRequest() *TaobaoUmpToolsGetAPIRequest {
	return poolTaobaoUmpToolsGetAPIRequest.Get().(*TaobaoUmpToolsGetAPIRequest)
}

// ReleaseTaobaoUmpToolsGetAPIRequest 将 TaobaoUmpToolsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpToolsGetAPIRequest(v *TaobaoUmpToolsGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpToolsGetAPIRequest.Put(v)
}
