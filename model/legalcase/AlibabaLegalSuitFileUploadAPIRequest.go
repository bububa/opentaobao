package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitFileUploadAPIRequest
诉讼文件上传接口 API请求
alibaba.legal.suit.file.upload

上传文件接口 */
type AlibabaLegalSuitFileUploadAPIRequest struct {
	model.Params
	// 文件
	_file *model.File
	// 时间搓
	_timeStamp int64
	// 文件名称
	_fileName string
	// 文件大小
	_fileSize int64
	// 签名
	_signature string
}

// New
