package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukumediaapivideosnapshotgetAPIRequest 根据视频ID查询视频缩微图 API请求
// youku.mediaapi.video.snapshot.get
//
// 根据视频ID查询视频缩微图
type YoukumediaapivideosnapshotgetAPIRequest struct {
	model.Params
	// 视频id
	_vid string
}

// NewYoukumediaapivideosnapshotgetRequest 初始化YoukumediaapivideosnapshotgetAPIRequest对象
func NewYoukumediaapivideosnapshotgetRequest() *YoukumediaapivideosnapshotgetAPIRequest {
	return &YoukumediaapivideosnapshotgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukumediaapivideosnapshotgetAPIRequest) GetApiMethodName() string {
	return "youku.mediaapi.video.snapshot.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukumediaapivideosnapshotgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukumediaapivideosnapshotgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVid is Vid Setter
// 视频id
func (r *YoukumediaapivideosnapshotgetAPIRequest) SetVid(_vid string) error {
	r._vid = _vid
	r.Set("vid", _vid)
	return nil
}

// GetVid Vid Getter
func (r YoukumediaapivideosnapshotgetAPIRequest) GetVid() string {
	return r._vid
}
