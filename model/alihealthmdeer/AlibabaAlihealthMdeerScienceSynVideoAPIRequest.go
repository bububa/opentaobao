package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmdeersciencesynVideoAPIRequest 视频同步【保存/更新】 API请求
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
type AlibabaalihealthmdeersciencesynVideoAPIRequest struct {
	model.Params
	// 视频信息实体
	_synVideoInfo *SynVideoInfo
}

// NewAlibabaalihealthmdeersciencesynVideoRequest 初始化AlibabaalihealthmdeersciencesynVideoAPIRequest对象
func NewAlibabaalihealthmdeersciencesynVideoRequest() *AlibabaalihealthmdeersciencesynVideoAPIRequest {
	return &AlibabaalihealthmdeersciencesynVideoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmdeersciencesynVideoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.synVideo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmdeersciencesynVideoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmdeersciencesynVideoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSynVideoInfo is SynVideoInfo Setter
// 视频信息实体
func (r *AlibabaalihealthmdeersciencesynVideoAPIRequest) SetSynVideoInfo(_synVideoInfo *SynVideoInfo) error {
	r._synVideoInfo = _synVideoInfo
	r.Set("syn_video_info", _synVideoInfo)
	return nil
}

// GetSynVideoInfo SynVideoInfo Getter
func (r AlibabaalihealthmdeersciencesynVideoAPIRequest) GetSynVideoInfo() *SynVideoInfo {
	return r._synVideoInfo
}
