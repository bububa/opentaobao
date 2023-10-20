package idleitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleItemMediaAddAPIRequest 图片上传 API请求
// alibaba.idle.item.media.add
//
// 上传图片
type AlibabaIdleItemMediaAddAPIRequest struct {
	model.Params
	// 废弃，不用再输入
	_userNick string
	// 多媒体文件字节流，图片<5M,视频<8M
	_mediaData *model.File
	// 类型：0 - 图片 ，仅支持图片
	_mediaType int64
}

// NewAlibabaIdleItemMediaAddRequest 初始化AlibabaIdleItemMediaAddAPIRequest对象
func NewAlibabaIdleItemMediaAddRequest() *AlibabaIdleItemMediaAddAPIRequest {
	return &AlibabaIdleItemMediaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleItemMediaAddAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.item.media.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleItemMediaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleItemMediaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 废弃，不用再输入
func (r *AlibabaIdleItemMediaAddAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabaIdleItemMediaAddAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetMediaData is MediaData Setter
// 多媒体文件字节流，图片&lt;5M,视频&lt;8M
func (r *AlibabaIdleItemMediaAddAPIRequest) SetMediaData(_mediaData *model.File) error {
	r._mediaData = _mediaData
	r.Set("media_data", _mediaData)
	return nil
}

// GetMediaData MediaData Getter
func (r AlibabaIdleItemMediaAddAPIRequest) GetMediaData() *model.File {
	return r._mediaData
}

// SetMediaType is MediaType Setter
// 类型：0 - 图片 ，仅支持图片
func (r *AlibabaIdleItemMediaAddAPIRequest) SetMediaType(_mediaType int64) error {
	r._mediaType = _mediaType
	r.Set("media_type", _mediaType)
	return nil
}

// GetMediaType MediaType Getter
func (r AlibabaIdleItemMediaAddAPIRequest) GetMediaType() int64 {
	return r._mediaType
}
