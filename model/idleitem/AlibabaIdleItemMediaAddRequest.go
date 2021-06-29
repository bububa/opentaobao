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
    _mediaData   *model.File
    // 类型：0 - 图片 ，仅支持图片
    _mediaType   int64
    // 废弃，不用再输入
    _userNick   string
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
func (r *AlibabaIdleItemMediaAddRequest) SetMediaData(_mediaData *model.File) error {
    r._mediaData = _mediaData
    r.Set("media_data", _mediaData)
    return nil
}

// MediaData Getter
func (r AlibabaIdleItemMediaAddRequest) GetMediaData() *model.File {
    return r._mediaData
}
// MediaType Setter
// 类型：0 - 图片 ，仅支持图片
func (r *AlibabaIdleItemMediaAddRequest) SetMediaType(_mediaType int64) error {
    r._mediaType = _mediaType
    r.Set("media_type", _mediaType)
    return nil
}

// MediaType Getter
func (r AlibabaIdleItemMediaAddRequest) GetMediaType() int64 {
    return r._mediaType
}
// UserNick Setter
// 废弃，不用再输入
func (r *AlibabaIdleItemMediaAddRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r AlibabaIdleItemMediaAddRequest) GetUserNick() string {
    return r._userNick
}
