package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemgameseverqueryAPIRequest 查询商品发布客户端下可用服务器列表 API请求
// alibaba.jym.item.game.sever.query
//
// 查询商品发布客户端下可用服务器列表
type AlibabajymitemgameseverqueryAPIRequest struct {
	model.Params
	// 游戏ID
	_gameId int64
	// 客户端ID
	_clientId int64
}

// NewAlibabajymitemgameseverqueryRequest 初始化AlibabajymitemgameseverqueryAPIRequest对象
func NewAlibabajymitemgameseverqueryRequest() *AlibabajymitemgameseverqueryAPIRequest {
	return &AlibabajymitemgameseverqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemgameseverqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.game.sever.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemgameseverqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemgameseverqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameId is GameId Setter
// 游戏ID
func (r *AlibabajymitemgameseverqueryAPIRequest) SetGameId(_gameId int64) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabajymitemgameseverqueryAPIRequest) GetGameId() int64 {
	return r._gameId
}

// SetClientId is ClientId Setter
// 客户端ID
func (r *AlibabajymitemgameseverqueryAPIRequest) SetClientId(_clientId int64) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabajymitemgameseverqueryAPIRequest) GetClientId() int64 {
	return r._clientId
}
