package wenyuvideo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuWenyuvideoSeetaGetAPIRequest 只看TA API请求
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
type YoukuWenyuvideoSeetaGetAPIRequest struct {
	model.Params
	// 视频字符串形式id
	_videoStrId string
}

// NewYoukuWenyuvideoSeetaGetRequest 初始化YoukuWenyuvideoSeetaGetAPIRequest对象
func NewYoukuWenyuvideoSeetaGetRequest() *YoukuWenyuvideoSeetaGetAPIRequest {
	return &YoukuWenyuvideoSeetaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuWenyuvideoSeetaGetAPIRequest) Reset() {
	r._videoStrId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoSeetaGetAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.seeta.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoSeetaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuWenyuvideoSeetaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVideoStrId is VideoStrId Setter
// 视频字符串形式id
func (r *YoukuWenyuvideoSeetaGetAPIRequest) SetVideoStrId(_videoStrId string) error {
	r._videoStrId = _videoStrId
	r.Set("video_str_id", _videoStrId)
	return nil
}

// GetVideoStrId VideoStrId Getter
func (r YoukuWenyuvideoSeetaGetAPIRequest) GetVideoStrId() string {
	return r._videoStrId
}

var poolYoukuWenyuvideoSeetaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuWenyuvideoSeetaGetRequest()
	},
}

// GetYoukuWenyuvideoSeetaGetRequest 从 sync.Pool 获取 YoukuWenyuvideoSeetaGetAPIRequest
func GetYoukuWenyuvideoSeetaGetAPIRequest() *YoukuWenyuvideoSeetaGetAPIRequest {
	return poolYoukuWenyuvideoSeetaGetAPIRequest.Get().(*YoukuWenyuvideoSeetaGetAPIRequest)
}

// ReleaseYoukuWenyuvideoSeetaGetAPIRequest 将 YoukuWenyuvideoSeetaGetAPIRequest 放入 sync.Pool
func ReleaseYoukuWenyuvideoSeetaGetAPIRequest(v *YoukuWenyuvideoSeetaGetAPIRequest) {
	v.Reset()
	poolYoukuWenyuvideoSeetaGetAPIRequest.Put(v)
}
