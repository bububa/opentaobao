package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiImagesegSegmentheadAPIRequest
头像分割 API请求
aliyun.viapi.imageseg.segmenthead

输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiImagesegSegmentheadAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New
