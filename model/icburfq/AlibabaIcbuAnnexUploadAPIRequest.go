package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuAnnexUploadAPIRequest
上传附件获取附件files_str API请求
alibaba.icbu.annex.upload

上传附件获取附件files_str */
type AlibabaIcbuAnnexUploadAPIRequest struct {
	model.Params
	// 文件名
	_fileName string
	// 文件字节流
	_fileInputStreamBytes *model.File
	// 来源
	_source string
}

// New
