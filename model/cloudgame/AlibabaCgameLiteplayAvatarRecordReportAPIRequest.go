package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCgameLiteplayAvatarRecordReportAPIRequest) Reset() {
	r._requestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.liteplay.avatar.record.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameLiteplayAvatarRecordReportAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCgameLiteplayAvatarRecordReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCgameLiteplayAvatarRecordReportRequest()
	},
}

// GetAlibabaCgameLiteplayAvatarRecordReportRequest 从 sync.Pool 获取 AlibabaCgameLiteplayAvatarRecordReportAPIRequest
func GetAlibabaCgameLiteplayAvatarRecordReportAPIRequest() *AlibabaCgameLiteplayAvatarRecordReportAPIRequest {
	return poolAlibabaCgameLiteplayAvatarRecordReportAPIRequest.Get().(*AlibabaCgameLiteplayAvatarRecordReportAPIRequest)
}

// ReleaseAlibabaCgameLiteplayAvatarRecordReportAPIRequest 将 AlibabaCgameLiteplayAvatarRecordReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaCgameLiteplayAvatarRecordReportAPIRequest(v *AlibabaCgameLiteplayAvatarRecordReportAPIRequest) {
	v.Reset()
	poolAlibabaCgameLiteplayAvatarRecordReportAPIRequest.Put(v)
}
