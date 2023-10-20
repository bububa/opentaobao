package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityAddAPIRequest 新增优惠活动 API请求
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
type TaobaoUmpActivityAddAPIRequest struct {
	model.Params
	// 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
	_content string
	// 工具id
	_toolId int64
}

// NewTaobaoUmpActivityAddRequest 初始化TaobaoUmpActivityAddAPIRequest对象
func NewTaobaoUmpActivityAddRequest() *TaobaoUmpActivityAddAPIRequest {
	return &TaobaoUmpActivityAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpActivityAddAPIRequest) Reset() {
	r._content = ""
	r._toolId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityAddAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpActivityAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
func (r *TaobaoUmpActivityAddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoUmpActivityAddAPIRequest) GetContent() string {
	return r._content
}

// SetToolId is ToolId Setter
// 工具id
func (r *TaobaoUmpActivityAddAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoUmpActivityAddAPIRequest) GetToolId() int64 {
	return r._toolId
}

var poolTaobaoUmpActivityAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpActivityAddRequest()
	},
}

// GetTaobaoUmpActivityAddRequest 从 sync.Pool 获取 TaobaoUmpActivityAddAPIRequest
func GetTaobaoUmpActivityAddAPIRequest() *TaobaoUmpActivityAddAPIRequest {
	return poolTaobaoUmpActivityAddAPIRequest.Get().(*TaobaoUmpActivityAddAPIRequest)
}

// ReleaseTaobaoUmpActivityAddAPIRequest 将 TaobaoUmpActivityAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpActivityAddAPIRequest(v *TaobaoUmpActivityAddAPIRequest) {
	v.Reset()
	poolTaobaoUmpActivityAddAPIRequest.Put(v)
}
