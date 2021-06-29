package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据视频ID查询视频缩微图 APIRequest
youku.mediaapi.video.snapshot.get

根据视频ID查询视频缩微图
*/
type YoukuMediaapiVideoSnapshotGetRequest struct {
    model.Params

    // 视频id
    vid   string 

}

func NewYoukuMediaapiVideoSnapshotGetRequest() *YoukuMediaapiVideoSnapshotGetRequest{
    return &YoukuMediaapiVideoSnapshotGetRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuMediaapiVideoSnapshotGetRequest) GetApiMethodName() string {
    return "youku.mediaapi.video.snapshot.get"
}

func (r YoukuMediaapiVideoSnapshotGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuMediaapiVideoSnapshotGetRequest) SetVid(vid string) error {
    r.vid = vid
    r.Set("vid", vid)
    return nil
}

func (r YoukuMediaapiVideoSnapshotGetRequest) GetVid() string {
    return r.vid
}

