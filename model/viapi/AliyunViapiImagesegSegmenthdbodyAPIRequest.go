package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiImagesegSegmenthdbodyAPIRequest
高清人体分割 API请求
aliyun.viapi.imageseg.segmenthdbody

对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiImagesegSegmenthdbodyAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New
