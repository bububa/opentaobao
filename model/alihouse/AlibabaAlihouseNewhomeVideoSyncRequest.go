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
type AlibabaAlihouseNewhomeVideoSyncAPIRequest struct {
    model.Params
    // 草稿对下
    _video   *VideoDraftDTO
}

// 初始化AlibabaAlihouseNewhomeVideoSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeVideoSyncRequest() *AlibabaAlihouseNewhomeVideoSyncAPIRequest{
    return &AlibabaAlihouseNewhomeVideoSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.video.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Video Setter
// 草稿对下
func (r *AlibabaAlihouseNewhomeVideoSyncAPIRequest) SetVideo(_video *VideoDraftDTO) error {
    r._video = _video
    r.Set("video", _video)
    return nil
}

// Video Getter
func (r AlibabaAlihouseNewhomeVideoSyncAPIRequest) GetVideo() *VideoDraftDTO {
    return r._video
}
