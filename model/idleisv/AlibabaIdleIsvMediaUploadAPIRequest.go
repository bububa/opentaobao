package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvMediaUploadAPIRequest
闲鱼服务商-图片上传 API请求
alibaba.idle.isv.media.upload

供外部服务商ISV进行闲鱼商品发布时上传商品所需图片 */
type AlibabaIdleIsvMediaUploadAPIRequest struct {
	model.Params
	// 多媒体字节数组
	_data *model.File
	// 文件名
	_name string
	// 0-表示图片，1-表示视频（暂不支持）
	_type int64
}

// New
