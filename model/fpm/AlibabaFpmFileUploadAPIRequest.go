package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFpmFileUploadAPIRequest
结算单文件上传 API请求
alibaba.fpm.file.upload

结算单文件上传 */
type AlibabaFpmFileUploadAPIRequest struct {
	model.Params
	// 实体
	_bizDto *FileUploadRequestDto
}

// New
