package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemGameSeverQueryAPIRequest 查询商品发布客户端下可用服务器列表 API请求
// alibaba.jym.item.game.sever.query
//
// 查询商品发布客户端下可用服务器列表
type AlibabaJymItemGameSeverQueryAPIRequest struct {
	model.Params
	// 游戏ID
	_gameId int64
	// 客户端ID
	_clientId int64
}

// NewAlibabaJymItemGameSeverQueryRequest 初始化AlibabaJymItemGameSeverQueryAPIRequest对象
func NewAlibabaJymItemGameSeverQueryRequest() *AlibabaJymItemGameSeverQueryAPIRequest {
	return &AlibabaJymItemGameSeverQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemGameSeverQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.game.sever.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemGameSeverQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemGameSeverQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameId is GameId Setter
// 游戏ID
func (r *AlibabaJymItemGameSeverQueryAPIRequest) SetGameId(_gameId int64) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabaJymItemGameSeverQueryAPIRequest) GetGameId() int64 {
	return r._gameId
}

// SetClientId is ClientId Setter
// 客户端ID
func (r *AlibabaJymItemGameSeverQueryAPIRequest) SetClientId(_clientId int64) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaJymItemGameSeverQueryAPIRequest) GetClientId() int64 {
	return r._clientId
}
