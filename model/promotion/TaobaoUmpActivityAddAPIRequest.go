package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivityaddAPIRequest 新增优惠活动 API请求
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
type TaobaoumpactivityaddAPIRequest struct {
	model.Params
	// 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
	_content string
	// 工具id
	_toolId int64
}

// NewTaobaoumpactivityaddRequest 初始化TaobaoumpactivityaddAPIRequest对象
func NewTaobaoumpactivityaddRequest() *TaobaoumpactivityaddAPIRequest {
	return &TaobaoumpactivityaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivityaddAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivityaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivityaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
func (r *TaobaoumpactivityaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoumpactivityaddAPIRequest) GetContent() string {
	return r._content
}

// SetToolId is ToolId Setter
// 工具id
func (r *TaobaoumpactivityaddAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoumpactivityaddAPIRequest) GetToolId() int64 {
	return r._toolId
}
