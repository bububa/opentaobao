package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsQueryAPIRequest 查询可配置任务素材接口 API请求
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
type TaobaoJstInteractiveAssetsQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveAssetsQueryRequest 初始化TaobaoJstInteractiveAssetsQueryAPIRequest对象
func NewTaobaoJstInteractiveAssetsQueryRequest() *TaobaoJstInteractiveAssetsQueryAPIRequest {
	return &TaobaoJstInteractiveAssetsQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveAssetsQueryAPIRequest) Reset() {
	r._miniAppId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsQueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveAssetsQueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

var poolTaobaoJstInteractiveAssetsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveAssetsQueryRequest()
	},
}

// GetTaobaoJstInteractiveAssetsQueryRequest 从 sync.Pool 获取 TaobaoJstInteractiveAssetsQueryAPIRequest
func GetTaobaoJstInteractiveAssetsQueryAPIRequest() *TaobaoJstInteractiveAssetsQueryAPIRequest {
	return poolTaobaoJstInteractiveAssetsQueryAPIRequest.Get().(*TaobaoJstInteractiveAssetsQueryAPIRequest)
}

// ReleaseTaobaoJstInteractiveAssetsQueryAPIRequest 将 TaobaoJstInteractiveAssetsQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveAssetsQueryAPIRequest(v *TaobaoJstInteractiveAssetsQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveAssetsQueryAPIRequest.Put(v)
}
