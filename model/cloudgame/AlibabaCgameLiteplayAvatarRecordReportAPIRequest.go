package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgameliteplayavatarrecordreportAPIRequest Avatar形象保存地址回调 API请求
// alibaba.cgame.liteplay.avatar.record.report
//
// 新氢玩, 围观互动Avatar捏脸, 形象地址保存回调
type AlibabacgameliteplayavatarrecordreportAPIRequest struct {
	model.Params
	// 请求消息体
	_requestDto *TopRecordCallbackRequest
}

// NewAlibabacgameliteplayavatarrecordreportRequest 初始化AlibabacgameliteplayavatarrecordreportAPIRequest对象
func NewAlibabacgameliteplayavatarrecordreportRequest() *AlibabacgameliteplayavatarrecordreportAPIRequest {
	return &AlibabacgameliteplayavatarrecordreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgameliteplayavatarrecordreportAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.liteplay.avatar.record.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgameliteplayavatarrecordreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgameliteplayavatarrecordreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestDto is RequestDto Setter
// 请求消息体
func (r *AlibabacgameliteplayavatarrecordreportAPIRequest) SetRequestDto(_requestDto *TopRecordCallbackRequest) error {
	r._requestDto = _requestDto
	r.Set("request_dto", _requestDto)
	return nil
}

// GetRequestDto RequestDto Getter
func (r AlibabacgameliteplayavatarrecordreportAPIRequest) GetRequestDto() *TopRecordCallbackRequest {
	return r._requestDto
}
