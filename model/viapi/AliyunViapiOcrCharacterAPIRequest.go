package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiOcrCharacterAPIRequest
通用文字识别 API请求
aliyun.viapi.ocr.character

获取通用的文字信息。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiOcrCharacterAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 图片类型, ,取范围[0:ImageURL, 1:ImageContent]
	_imageType int64
	// 图片中文字的最小高度，单位像素
	_minHeight int64
	// 是否输出文字框的概率,取值范围[true:是, false:否]
	_outputProbability bool
}

// New
