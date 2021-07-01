package idleitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleItemMediaAddAPIRequest
图片上传 API请求
alibaba.idle.item.media.add

上传图片 */
type AlibabaIdleItemMediaAddAPIRequest struct {
	model.Params
	// 多媒体文件字节流，图片<5M,视频<8M
	_mediaData *model.File
	// 类型：0 - 图片 ，仅支持图片
	_mediaType int64
	// 废弃，不用再输入
	_userNick string
}

// New
