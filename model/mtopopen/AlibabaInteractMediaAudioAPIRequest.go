package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractMediaAudioAPIRequest 音频相关鉴权接口 API请求
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
type AlibabaInteractMediaAudioAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractMediaAudioRequest 初始化AlibabaInteractMediaAudioAPIRequest对象
func NewAlibabaInteractMediaAudioRequest() *AlibabaInteractMediaAudioAPIRequest {
	return &AlibabaInteractMediaAudioAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractMediaAudioAPIRequest) Reset() {
	r._id = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaAudioAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.audio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaAudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractMediaAudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaAudioAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractMediaAudioAPIRequest) GetId() string {
	return r._id
}

var poolAlibabaInteractMediaAudioAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractMediaAudioRequest()
	},
}

// GetAlibabaInteractMediaAudioRequest 从 sync.Pool 获取 AlibabaInteractMediaAudioAPIRequest
func GetAlibabaInteractMediaAudioAPIRequest() *AlibabaInteractMediaAudioAPIRequest {
	return poolAlibabaInteractMediaAudioAPIRequest.Get().(*AlibabaInteractMediaAudioAPIRequest)
}

// ReleaseAlibabaInteractMediaAudioAPIRequest 将 AlibabaInteractMediaAudioAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractMediaAudioAPIRequest(v *AlibabaInteractMediaAudioAPIRequest) {
	v.Reset()
	poolAlibabaInteractMediaAudioAPIRequest.Put(v)
}
