package jym

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest 基于游戏id临时上下架智能发布入口 API请求
// alibaba.jym.goods.external.goods.vmos.offon.game
//
// 基于游戏id临时上下架智能发布入口
type AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest struct {
	model.Params
	// 所需要下架的游戏ID值
	_gameId int64
	// offGame是否下架游戏，true代表下架，false代表上架
	_offGame bool
}

// NewAlibabaJymGoodsExternalGoodsVmosOffonGameRequest 初始化AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest对象
func NewAlibabaJymGoodsExternalGoodsVmosOffonGameRequest() *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest {
	return &AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) Reset() {
	r._gameId = 0
	r._offGame = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.goods.external.goods.vmos.offon.game"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameId is GameId Setter
// 所需要下架的游戏ID值
func (r *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) SetGameId(_gameId int64) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) GetGameId() int64 {
	return r._gameId
}

// SetOffGame is OffGame Setter
// offGame是否下架游戏，true代表下架，false代表上架
func (r *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) SetOffGame(_offGame bool) error {
	r._offGame = _offGame
	r.Set("off_game", _offGame)
	return nil
}

// GetOffGame OffGame Getter
func (r AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) GetOffGame() bool {
	return r._offGame
}

var poolAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymGoodsExternalGoodsVmosOffonGameRequest()
	},
}

// GetAlibabaJymGoodsExternalGoodsVmosOffonGameRequest 从 sync.Pool 获取 AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest
func GetAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest() *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest {
	return poolAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest.Get().(*AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest)
}

// ReleaseAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest 将 AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest(v *AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest) {
	v.Reset()
	poolAlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest.Put(v)
}
