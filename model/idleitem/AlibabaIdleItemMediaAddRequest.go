package idleitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传 API请求
alibaba.idle.item.media.add

上传图片
*/
type AlibabaIdleItemMediaAddRequest struct {
    model.Params
    // 多媒体文件字节流，图片<5M,视频<8M
    mediaData   []*model.File
    // 类型：0 - 图片 ，仅支持图片
    mediaType   int64
    // 废弃，不用再输入
    userNick   string
}

// 初始化AlibabaIdleItemMediaAddRequest对象
func NewAlibabaIdleItemMediaAddRequest() *AlibabaIdleItemMediaAddRequest{
    return &AlibabaIdleItemMediaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemMediaAddRequest) GetApiMethodName() string {
    return "alibaba.idle.item.media.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemMediaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MediaData Setter
// 多媒体文件字节流，图片<5M,视频<8M
func (r *AlibabaIdleItemMediaAddRequest) SetMediaData(mediaData []*model.File) error {
    r.mediaData = mediaData
    r.Set("media_data", mediaData)
    return nil
}

// MediaData Getter
func (r AlibabaIdleItemMediaAddRequest) GetMediaData() []*model.File {
    return r.mediaData
}
// MediaType Setter
// 类型：0 - 图片 ，仅支持图片
func (r *AlibabaIdleItemMediaAddRequest) SetMediaType(mediaType int64) error {
    r.mediaType = mediaType
    r.Set("media_type", mediaType)
    return nil
}

// MediaType Getter
func (r AlibabaIdleItemMediaAddRequest) GetMediaType() int64 {
    return r.mediaType
}
// UserNick Setter
// 废弃，不用再输入
func (r *AlibabaIdleItemMediaAddRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r AlibabaIdleItemMediaAddRequest) GetUserNick() string {
    return r.userNick
}
