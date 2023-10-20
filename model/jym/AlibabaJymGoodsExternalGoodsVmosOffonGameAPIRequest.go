package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest 基于游戏id临时上下架智能发布入口 API请求
// alibaba.jym.goods.external.goods.vmos.offon.game
//
// 基于游戏id临时上下架智能发布入口
type AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest struct {
	model.Params
	// 所需要下架的游戏ID值
	_gameId int64
	// offGame是否下架游戏，true代表下架，false代表上架
	_offGame bool
}

// NewAlibabajymgoodsexternalgoodsvmosoffongameRequest 初始化AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest对象
func NewAlibabajymgoodsexternalgoodsvmosoffongameRequest() *AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest {
	return &AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.goods.external.goods.vmos.offon.game"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameId is GameId Setter
// 所需要下架的游戏ID值
func (r *AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) SetGameId(_gameId int64) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) GetGameId() int64 {
	return r._gameId
}

// SetOffGame is OffGame Setter
// offGame是否下架游戏，true代表下架，false代表上架
func (r *AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) SetOffGame(_offGame bool) error {
	r._offGame = _offGame
	r.Set("off_game", _offGame)
	return nil
}

// GetOffGame OffGame Getter
func (r AlibabajymgoodsexternalgoodsvmosoffongameAPIRequest) GetOffGame() bool {
	return r._offGame
}
