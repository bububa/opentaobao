package wenyuvideo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
只看TA API请求
youku.wenyuvideo.seeta.get

只看Ta对外输出
*/
type YoukuWenyuvideoSeetaGetRequest struct {
    model.Params
    // 视频字符串形式id
    _videoStrId   string
}

// 初始化YoukuWenyuvideoSeetaGetRequest对象
func NewYoukuWenyuvideoSeetaGetRequest() *YoukuWenyuvideoSeetaGetRequest{
    return &YoukuWenyuvideoSeetaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoSeetaGetRequest) GetApiMethodName() string {
    return "youku.wenyuvideo.seeta.get"
}

// IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoSeetaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VideoStrId Setter
// 视频字符串形式id
func (r *YoukuWenyuvideoSeetaGetRequest) SetVideoStrId(_videoStrId string) error {
    r._videoStrId = _videoStrId
    r.Set("video_str_id", _videoStrId)
    return nil
}

// VideoStrId Getter
func (r YoukuWenyuvideoSeetaGetRequest) GetVideoStrId() string {
    return r._videoStrId
}
