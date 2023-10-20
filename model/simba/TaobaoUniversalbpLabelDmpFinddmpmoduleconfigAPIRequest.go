package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest 查询dmp浮层配置 API请求
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
type TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpLabelDmpFinddmpmoduleconfigRequest 初始化TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest对象
func NewTaobaoUniversalbpLabelDmpFinddmpmoduleconfigRequest() *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest {
	return &TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) Reset() {
	r._topServiceContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.label.dmp.finddmpmoduleconfig"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

var poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpLabelDmpFinddmpmoduleconfigRequest()
	},
}

// GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigRequest 从 sync.Pool 获取 TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest
func GetTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest() *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest {
	return poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest.Get().(*TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest)
}

// ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest 将 TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest(v *TaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpLabelDmpFinddmpmoduleconfigAPIRequest.Put(v)
}
