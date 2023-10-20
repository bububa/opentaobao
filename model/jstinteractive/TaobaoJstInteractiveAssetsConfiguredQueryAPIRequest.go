package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest 查询已配置的任务素材列表接口 API请求
// taobao.jst.interactive.assets.configured.query
//
// 查询已配置任务素材列表
type TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// NewTaobaoJstInteractiveAssetsConfiguredQueryRequest 初始化TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest对象
func NewTaobaoJstInteractiveAssetsConfiguredQueryRequest() *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest {
	return &TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) Reset() {
	r._miniAppId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.assets.configured.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

var poolTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveAssetsConfiguredQueryRequest()
	},
}

// GetTaobaoJstInteractiveAssetsConfiguredQueryRequest 从 sync.Pool 获取 TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest
func GetTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest() *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest {
	return poolTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest.Get().(*TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest)
}

// ReleaseTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest 将 TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest(v *TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveAssetsConfiguredQueryAPIRequest.Put(v)
}
