package lstspeacker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstSpeakerFileUploadAPIRequest
如意音箱音频文件长传 API请求
alibaba.lst.speaker.file.upload

如意音箱音频文件长传 */
type AlibabaLstSpeakerFileUploadAPIRequest struct {
	model.Params
	// 数据流
	_fileBytes *model.File
	// 文件类型,audio:音频，advert:广告
	_fileType string
	// 文件ID
	_fileId string
	// md5直
	_md5 string
}

// New
