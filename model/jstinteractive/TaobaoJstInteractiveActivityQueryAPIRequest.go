package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveActivityQueryAPIRequest 互动任务活动查询接口 API请求
// taobao.jst.interactive.activity.query
//
// 互动任务活动查询接口
type TaobaoJstInteractiveActivityQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveActivityQueryRequest 初始化TaobaoJstInteractiveActivityQueryAPIRequest对象
func NewTaobaoJstInteractiveActivityQueryRequest() *TaobaoJstInteractiveActivityQueryAPIRequest {
	return &TaobaoJstInteractiveActivityQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveActivityQueryAPIRequest) Reset() {
	r._miniAppId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityQueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveActivityQueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

var poolTaobaoJstInteractiveActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveActivityQueryRequest()
	},
}

// GetTaobaoJstInteractiveActivityQueryRequest 从 sync.Pool 获取 TaobaoJstInteractiveActivityQueryAPIRequest
func GetTaobaoJstInteractiveActivityQueryAPIRequest() *TaobaoJstInteractiveActivityQueryAPIRequest {
	return poolTaobaoJstInteractiveActivityQueryAPIRequest.Get().(*TaobaoJstInteractiveActivityQueryAPIRequest)
}

// ReleaseTaobaoJstInteractiveActivityQueryAPIRequest 将 TaobaoJstInteractiveActivityQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveActivityQueryAPIRequest(v *TaobaoJstInteractiveActivityQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveActivityQueryAPIRequest.Put(v)
}
