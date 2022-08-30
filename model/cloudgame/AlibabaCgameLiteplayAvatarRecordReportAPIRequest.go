package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameLiteplayAvatarRecordReportAPIRequest Avatar形象保存地址回调 API请求
// alibaba.cgame.liteplay.avatar.record.report
//
// 新氢玩, 围观互动Avatar捏脸, 形象地址保存回调
type AlibabaCgameLiteplayAvatarRecordReportAPIRequest struct {
	model.Params
	// 请求消息体
	_requestDto *TopRecordCallbackRequest
}

// NewAlibabaCgameLiteplayAvatarRecordReportRequest 初始化AlibabaCgameLiteplayAvatarRecordReportAPIRequest对象
func NewAlibabaCgameLiteplayAvatarRecordReportRequest() *AlibabaCgameLiteplayAvatarRecordReportAPIRequest {
	return &AlibabaCgameLiteplayAvatarRecordReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.liteplay.avatar.record.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequestDto is RequestDto Setter
// 请求消息体
func (r *AlibabaCgameLiteplayAvatarRecordReportAPIRequest) SetRequestDto(_requestDto *TopRecordCallbackRequest) error {
	r._requestDto = _requestDto
	r.Set("request_dto", _requestDto)
	return nil
}

// GetRequestDto RequestDto Getter
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetRequestDto() *TopRecordCallbackRequest {
	return r._requestDto
}
