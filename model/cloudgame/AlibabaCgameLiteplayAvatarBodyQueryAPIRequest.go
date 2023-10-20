package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgameliteplayavatarbodyqueryAPIRequest 新氢玩Avatar脸部装扮数据查询 API请求
// alibaba.cgame.liteplay.avatar.body.query
//
// 云游戏, 新氢玩, 围观互动,自研游戏包, 需要查询Avatar虚拟形象捏脸和装扮数据, 用来初始化游戏包形象.
type AlibabacgameliteplayavatarbodyqueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestDto *TopQueryUserBodyDressRequest
}

// NewAlibabacgameliteplayavatarbodyqueryRequest 初始化AlibabacgameliteplayavatarbodyqueryAPIRequest对象
func NewAlibabacgameliteplayavatarbodyqueryRequest() *AlibabacgameliteplayavatarbodyqueryAPIRequest {
	return &AlibabacgameliteplayavatarbodyqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgameliteplayavatarbodyqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.liteplay.avatar.body.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgameliteplayavatarbodyqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgameliteplayavatarbodyqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestDto is RequestDto Setter
// 请求入参
func (r *AlibabacgameliteplayavatarbodyqueryAPIRequest) SetRequestDto(_requestDto *TopQueryUserBodyDressRequest) error {
	r._requestDto = _requestDto
	r.Set("request_dto", _requestDto)
	return nil
}

// GetRequestDto RequestDto Getter
func (r AlibabacgameliteplayavatarbodyqueryAPIRequest) GetRequestDto() *TopQueryUserBodyDressRequest {
	return r._requestDto
}
