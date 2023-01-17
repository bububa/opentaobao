package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest 本地生活用户增长互动小游戏积分发放 API请求
// alibaba.alsc.growth.interactive.mini.game.integral.grant
//
// 本地生活用户增长互动小游戏积分发放
type AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest struct {
	model.Params
	// 小游戏积分发放请求参数
	_miniGameGrantIntegralRequest *MiniGameGrantIntegralRequest
}

// NewAlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest 初始化AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest对象
func NewAlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest() *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest {
	return &AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.mini.game.integral.grant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniGameGrantIntegralRequest is MiniGameGrantIntegralRequest Setter
// 小游戏积分发放请求参数
func (r *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest) SetMiniGameGrantIntegralRequest(_miniGameGrantIntegralRequest *MiniGameGrantIntegralRequest) error {
	r._miniGameGrantIntegralRequest = _miniGameGrantIntegralRequest
	r.Set("mini_game_grant_integral_request", _miniGameGrantIntegralRequest)
	return nil
}

// GetMiniGameGrantIntegralRequest MiniGameGrantIntegralRequest Getter
func (r AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest) GetMiniGameGrantIntegralRequest() *MiniGameGrantIntegralRequest {
	return r._miniGameGrantIntegralRequest
}
