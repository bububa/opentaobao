package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameLiteplayAvatarBodyQueryAPIRequest 新氢玩Avatar脸部装扮数据查询 API请求
// alibaba.cgame.liteplay.avatar.body.query
//
// 云游戏, 新氢玩, 围观互动,自研游戏包, 需要查询Avatar虚拟形象捏脸和装扮数据, 用来初始化游戏包形象.
type AlibabaCgameLiteplayAvatarBodyQueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestDto *TopQueryUserBodyDressRequest
}

// NewAlibabaCgameLiteplayAvatarBodyQueryRequest 初始化AlibabaCgameLiteplayAvatarBodyQueryAPIRequest对象
func NewAlibabaCgameLiteplayAvatarBodyQueryRequest() *AlibabaCgameLiteplayAvatarBodyQueryAPIRequest {
	return &AlibabaCgameLiteplayAvatarBodyQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameLiteplayAvatarBodyQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.liteplay.avatar.body.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameLiteplayAvatarBodyQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequestDto is RequestDto Setter
// 请求入参
func (r *AlibabaCgameLiteplayAvatarBodyQueryAPIRequest) SetRequestDto(_requestDto *TopQueryUserBodyDressRequest) error {
	r._requestDto = _requestDto
	r.Set("request_dto", _requestDto)
	return nil
}

// GetRequestDto RequestDto Getter
func (r AlibabaCgameLiteplayAvatarBodyQueryAPIRequest) GetRequestDto() *TopQueryUserBodyDressRequest {
	return r._requestDto
}
