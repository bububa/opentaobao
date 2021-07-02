package youkuott

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetApiMethodName() string {
	return "youku.mediaapi.video.snapshot.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Vid Setter
// 视频id
func (r *YoukuMediaapiVideoSnapshotGetAPIRequest) SetVid(_vid string) error {
	r._vid = _vid
	r.Set("vid", _vid)
	return nil
}

// Get Vid Getter
func (r YoukuMediaapiVideoSnapshotGetAPIRequest) GetVid() string {
	return r._vid
}
