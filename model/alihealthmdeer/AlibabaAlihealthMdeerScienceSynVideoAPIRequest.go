package alihealthmdeer

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceSynVideoAPIRequest 视频同步【保存/更新】 API请求
// alibaba.alihealth.mdeer.science.synVideo
//
// 视频同步【保存/更新】
type AlibabaAlihealthMdeerScienceSynVideoAPIRequest struct {
	model.Params
	// 视频信息实体
	_synVideoInfo *SynVideoInfo
}

// NewAlibabaAlihealthMdeerScienceSynVideoRequest 初始化AlibabaAlihealthMdeerScienceSynVideoAPIRequest对象
func NewAlibabaAlihealthMdeerScienceSynVideoRequest() *AlibabaAlihealthMdeerScienceSynVideoAPIRequest {
	return &AlibabaAlihealthMdeerScienceSynVideoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMdeerScienceSynVideoAPIRequest) Reset() {
	r._synVideoInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.science.synVideo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSynVideoInfo is SynVideoInfo Setter
// 视频信息实体
func (r *AlibabaAlihealthMdeerScienceSynVideoAPIRequest) SetSynVideoInfo(_synVideoInfo *SynVideoInfo) error {
	r._synVideoInfo = _synVideoInfo
	r.Set("syn_video_info", _synVideoInfo)
	return nil
}

// GetSynVideoInfo SynVideoInfo Getter
func (r AlibabaAlihealthMdeerScienceSynVideoAPIRequest) GetSynVideoInfo() *SynVideoInfo {
	return r._synVideoInfo
}

var poolAlibabaAlihealthMdeerScienceSynVideoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMdeerScienceSynVideoRequest()
	},
}

// GetAlibabaAlihealthMdeerScienceSynVideoRequest 从 sync.Pool 获取 AlibabaAlihealthMdeerScienceSynVideoAPIRequest
func GetAlibabaAlihealthMdeerScienceSynVideoAPIRequest() *AlibabaAlihealthMdeerScienceSynVideoAPIRequest {
	return poolAlibabaAlihealthMdeerScienceSynVideoAPIRequest.Get().(*AlibabaAlihealthMdeerScienceSynVideoAPIRequest)
}

// ReleaseAlibabaAlihealthMdeerScienceSynVideoAPIRequest 将 AlibabaAlihealthMdeerScienceSynVideoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMdeerScienceSynVideoAPIRequest(v *AlibabaAlihealthMdeerScienceSynVideoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMdeerScienceSynVideoAPIRequest.Put(v)
}
