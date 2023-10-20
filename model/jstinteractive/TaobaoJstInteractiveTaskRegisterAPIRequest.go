package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveTaskRegisterAPIRequest 互动任务开通接口 API请求
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
type TaobaoJstInteractiveTaskRegisterAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveTaskRegisterRequest 初始化TaobaoJstInteractiveTaskRegisterAPIRequest对象
func NewTaobaoJstInteractiveTaskRegisterRequest() *TaobaoJstInteractiveTaskRegisterAPIRequest {
	return &TaobaoJstInteractiveTaskRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveTaskRegisterAPIRequest) Reset() {
	r._miniAppId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveTaskRegisterAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveTaskRegisterAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

var poolTaobaoJstInteractiveTaskRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveTaskRegisterRequest()
	},
}

// GetTaobaoJstInteractiveTaskRegisterRequest 从 sync.Pool 获取 TaobaoJstInteractiveTaskRegisterAPIRequest
func GetTaobaoJstInteractiveTaskRegisterAPIRequest() *TaobaoJstInteractiveTaskRegisterAPIRequest {
	return poolTaobaoJstInteractiveTaskRegisterAPIRequest.Get().(*TaobaoJstInteractiveTaskRegisterAPIRequest)
}

// ReleaseTaobaoJstInteractiveTaskRegisterAPIRequest 将 TaobaoJstInteractiveTaskRegisterAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveTaskRegisterAPIRequest(v *TaobaoJstInteractiveTaskRegisterAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveTaskRegisterAPIRequest.Put(v)
}
