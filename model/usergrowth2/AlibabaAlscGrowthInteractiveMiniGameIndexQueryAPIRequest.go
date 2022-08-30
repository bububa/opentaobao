package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest 小游戏首页查询接口 API请求
// alibaba.alsc.growth.interactive.mini.game.index.query
//
// 小游戏首页查询接口
type AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest struct {
	model.Params
	// 小游戏首页请求对象
	_miniGameIndexRequest *MiniGameIndexRequest
}

// NewAlibabaAlscGrowthInteractiveMiniGameIndexQueryRequest 初始化AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest对象
func NewAlibabaAlscGrowthInteractiveMiniGameIndexQueryRequest() *AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest {
	return &AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.mini.game.index.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMiniGameIndexRequest is MiniGameIndexRequest Setter
// 小游戏首页请求对象
func (r *AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest) SetMiniGameIndexRequest(_miniGameIndexRequest *MiniGameIndexRequest) error {
	r._miniGameIndexRequest = _miniGameIndexRequest
	r.Set("mini_game_index_request", _miniGameIndexRequest)
	return nil
}

// GetMiniGameIndexRequest MiniGameIndexRequest Getter
func (r AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIRequest) GetMiniGameIndexRequest() *MiniGameIndexRequest {
	return r._miniGameIndexRequest
}
