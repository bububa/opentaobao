package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalImDataUploadAPIRequest
三方IM图片音频消息上传 API请求
alibaba.alihealth.medical.im.data.upload

三方IM图片音频消息上传 */
type AlibabaAlihealthMedicalImDataUploadAPIRequest struct {
	model.Params
	// request
	_uploadDataRequest *UploadDataRequest
	// 文件字节流
	_file *model.File
}

// New
