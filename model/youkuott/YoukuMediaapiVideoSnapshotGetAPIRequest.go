package youkuott

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuMediaapiVideoSnapshotGetAPIRequest 根据视频ID查询视频缩微图 API请求
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
type YoukuMediaapiVideoSnapshotGetAPIRequest struct {
	model.Params
	// 视频id
	_vid string
}

// NewYoukuMediaapiVideoSnapshotGetRequest 初始化YoukuMediaapiVideoSnapshotGetAPIRequest对象
func NewYoukuMediaapiVideoSnapshotGetRequest() *YoukuMediaapiVideoSnapshotGetAPIRequest {
	return &YoukuMediaapiVideoSnapshotGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuMediaapiVideoSnapshotGetAPIRequest) Reset() {
	r._vid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetApiMethodName() string {
	return "youku.mediaapi.video.snapshot.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVid is Vid Setter
// 视频id
func (r *YoukuMediaapiVideoSnapshotGetAPIRequest) SetVid(_vid string) error {
	r._vid = _vid
	r.Set("vid", _vid)
	return nil
}

// GetVid Vid Getter
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetVid() string {
	return r._vid
}

var poolYoukuMediaapiVideoSnapshotGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuMediaapiVideoSnapshotGetRequest()
	},
}

// GetYoukuMediaapiVideoSnapshotGetRequest 从 sync.Pool 获取 YoukuMediaapiVideoSnapshotGetAPIRequest
func GetYoukuMediaapiVideoSnapshotGetAPIRequest() *YoukuMediaapiVideoSnapshotGetAPIRequest {
	return poolYoukuMediaapiVideoSnapshotGetAPIRequest.Get().(*YoukuMediaapiVideoSnapshotGetAPIRequest)
}

// ReleaseYoukuMediaapiVideoSnapshotGetAPIRequest 将 YoukuMediaapiVideoSnapshotGetAPIRequest 放入 sync.Pool
func ReleaseYoukuMediaapiVideoSnapshotGetAPIRequest(v *YoukuMediaapiVideoSnapshotGetAPIRequest) {
	v.Reset()
	poolYoukuMediaapiVideoSnapshotGetAPIRequest.Put(v)
}
