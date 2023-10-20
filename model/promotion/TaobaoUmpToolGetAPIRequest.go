package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpToolGetAPIRequest) Reset() {
	r._toolId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.tool.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpToolGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToolId is ToolId Setter
// 工具的id
func (r *TaobaoUmpToolGetAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoUmpToolGetAPIRequest) GetToolId() int64 {
	return r._toolId
}

var poolTaobaoUmpToolGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpToolGetRequest()
	},
}

// GetTaobaoUmpToolGetRequest 从 sync.Pool 获取 TaobaoUmpToolGetAPIRequest
func GetTaobaoUmpToolGetAPIRequest() *TaobaoUmpToolGetAPIRequest {
	return poolTaobaoUmpToolGetAPIRequest.Get().(*TaobaoUmpToolGetAPIRequest)
}

// ReleaseTaobaoUmpToolGetAPIRequest 将 TaobaoUmpToolGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpToolGetAPIRequest(v *TaobaoUmpToolGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpToolGetAPIRequest.Put(v)
}
