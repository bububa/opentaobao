package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiImagesegSegmentcomodityAPIRequest
商品分割 API请求
aliyun.viapi.imageseg.segmentcomodity

识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiImagesegSegmentcomodityAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New
