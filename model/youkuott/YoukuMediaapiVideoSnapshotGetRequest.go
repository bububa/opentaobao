package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据视频ID查询视频缩微图 API请求
youku.mediaapi.video.snapshot.get

根据视频ID查询视频缩微图
*/
type YoukuMediaapiVideoSnapshotGetRequest struct {
    model.Params
    // 视频id
    vid   string
}

// 初始化YoukuMediaapiVideoSnapshotGetRequest对象
func NewYoukuMediaapiVideoSnapshotGetRequest() *YoukuMediaapiVideoSnapshotGetRequest{
    return &YoukuMediaapiVideoSnapshotGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuMediaapiVideoSnapshotGetRequest) GetApiMethodName() string {
    return "youku.mediaapi.video.snapshot.get"
}

// IRequest interface 方法, 获取API参数
func (r YoukuMediaapiVideoSnapshotGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vid Setter
// 视频id
func (r *YoukuMediaapiVideoSnapshotGetRequest) SetVid(vid string) error {
    r.vid = vid
    r.Set("vid", vid)
    return nil
}

// Vid Getter
func (r YoukuMediaapiVideoSnapshotGetRequest) GetVid() string {
    return r.vid
}
