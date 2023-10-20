package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdzoneFindconfiglistAPIRequest 查询所有可用资源包信息 API请求
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
type TaobaoUniversalbpAdzoneFindconfiglistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpAdzoneFindconfiglistRequest 初始化TaobaoUniversalbpAdzoneFindconfiglistAPIRequest对象
func NewTaobaoUniversalbpAdzoneFindconfiglistRequest() *TaobaoUniversalbpAdzoneFindconfiglistAPIRequest {
	return &TaobaoUniversalbpAdzoneFindconfiglistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) Reset() {
	r._topServiceContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.adzone.findconfiglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

var poolTaobaoUniversalbpAdzoneFindconfiglistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpAdzoneFindconfiglistRequest()
	},
}

// GetTaobaoUniversalbpAdzoneFindconfiglistRequest 从 sync.Pool 获取 TaobaoUniversalbpAdzoneFindconfiglistAPIRequest
func GetTaobaoUniversalbpAdzoneFindconfiglistAPIRequest() *TaobaoUniversalbpAdzoneFindconfiglistAPIRequest {
	return poolTaobaoUniversalbpAdzoneFindconfiglistAPIRequest.Get().(*TaobaoUniversalbpAdzoneFindconfiglistAPIRequest)
}

// ReleaseTaobaoUniversalbpAdzoneFindconfiglistAPIRequest 将 TaobaoUniversalbpAdzoneFindconfiglistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpAdzoneFindconfiglistAPIRequest(v *TaobaoUniversalbpAdzoneFindconfiglistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpAdzoneFindconfiglistAPIRequest.Put(v)
}
