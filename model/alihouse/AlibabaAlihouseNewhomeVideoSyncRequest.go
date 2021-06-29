package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频草稿信息同步 API请求
alibaba.alihouse.newhome.video.sync

接收视频信息记录
*/
type AlibabaAlihouseNewhomeVideoSyncRequest struct {
    model.Params
    // 草稿对下
    _video   *VideoDraftDTO
}

// 初始化AlibabaAlihouseNewhomeVideoSyncRequest对象
func NewAlibabaAlihouseNewhomeVideoSyncRequest() *AlibabaAlihouseNewhomeVideoSyncRequest{
    return &AlibabaAlihouseNewhomeVideoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.video.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Video Setter
// 草稿对下
func (r *AlibabaAlihouseNewhomeVideoSyncRequest) SetVideo(_video *VideoDraftDTO) error {
    r._video = _video
    r.Set("video", _video)
    return nil
}

// Video Getter
func (r AlibabaAlihouseNewhomeVideoSyncRequest) GetVideo() *VideoDraftDTO {
    return r._video
}
