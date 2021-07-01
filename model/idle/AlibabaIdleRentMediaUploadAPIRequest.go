package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentMediaUploadAPIRequest
闲鱼多媒体上传接口 API请求
alibaba.idle.rent.media.upload

上传多媒体信息，包括图片、视频（暂不支持） */
type AlibabaIdleRentMediaUploadAPIRequest struct {
	model.Params
	// 多媒体字节数组
	_data *model.File
	// 文件名
	_name string
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// New
