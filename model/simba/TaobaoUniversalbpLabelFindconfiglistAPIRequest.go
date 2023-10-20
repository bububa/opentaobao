package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpLabelFindconfiglistAPIRequest 查询可用标签id信息 API请求
// taobao.universalbp.label.findconfiglist
//
// 入参账号信息，出参可用标签id，用于下游接口入参
type TaobaoUniversalbpLabelFindconfiglistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpLabelFindconfiglistRequest 初始化TaobaoUniversalbpLabelFindconfiglistAPIRequest对象
func NewTaobaoUniversalbpLabelFindconfiglistRequest() *TaobaoUniversalbpLabelFindconfiglistAPIRequest {
	return &TaobaoUniversalbpLabelFindconfiglistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpLabelFindconfiglistAPIRequest) Reset() {
	r._topServiceContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpLabelFindconfiglistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.label.findconfiglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpLabelFindconfiglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpLabelFindconfiglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpLabelFindconfiglistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpLabelFindconfiglistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

var poolTaobaoUniversalbpLabelFindconfiglistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpLabelFindconfiglistRequest()
	},
}

// GetTaobaoUniversalbpLabelFindconfiglistRequest 从 sync.Pool 获取 TaobaoUniversalbpLabelFindconfiglistAPIRequest
func GetTaobaoUniversalbpLabelFindconfiglistAPIRequest() *TaobaoUniversalbpLabelFindconfiglistAPIRequest {
	return poolTaobaoUniversalbpLabelFindconfiglistAPIRequest.Get().(*TaobaoUniversalbpLabelFindconfiglistAPIRequest)
}

// ReleaseTaobaoUniversalbpLabelFindconfiglistAPIRequest 将 TaobaoUniversalbpLabelFindconfiglistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpLabelFindconfiglistAPIRequest(v *TaobaoUniversalbpLabelFindconfiglistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpLabelFindconfiglistAPIRequest.Put(v)
}
