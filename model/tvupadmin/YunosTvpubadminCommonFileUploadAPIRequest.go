package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminCommonFileUploadAPIRequest
文件上传API API请求
yunos.tvpubadmin.common.file.upload

文件上传服务 */
type YunosTvpubadminCommonFileUploadAPIRequest struct {
	model.Params
	// 文件字节流
	_bytes *model.File
	// 原文件名
	_originalFilename string
	// 文件大小
	_size string
	// 文件类型
	_contentType string
	// 上传地址
	_uploadPath string
}

// New
