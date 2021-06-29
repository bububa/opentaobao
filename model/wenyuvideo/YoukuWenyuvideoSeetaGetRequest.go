package wenyuvideo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
只看TA APIRequest
youku.wenyuvideo.seeta.get

只看Ta对外输出
*/
type YoukuWenyuvideoSeetaGetRequest struct {
    model.Params

    // 视频字符串形式id
    videoStrId   string 

}

func NewYoukuWenyuvideoSeetaGetRequest() *YoukuWenyuvideoSeetaGetRequest{
    return &YoukuWenyuvideoSeetaGetRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuWenyuvideoSeetaGetRequest) GetApiMethodName() string {
    return "youku.wenyuvideo.seeta.get"
}

func (r YoukuWenyuvideoSeetaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuWenyuvideoSeetaGetRequest) SetVideoStrId(videoStrId string) error {
    r.videoStrId = videoStrId
    r.Set("video_str_id", videoStrId)
    return nil
}

func (r YoukuWenyuvideoSeetaGetRequest) GetVideoStrId() string {
    return r.videoStrId
}

