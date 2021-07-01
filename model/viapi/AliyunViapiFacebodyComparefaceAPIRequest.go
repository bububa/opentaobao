package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiFacebodyComparefaceAPIRequest
人脸比对1：1 API请求
aliyun.viapi.facebody.compareface

基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiFacebodyComparefaceAPIRequest struct {
	model.Params
	// 图片url地址(http/https)
	_imageUrlA string
	// 图片url地址(http/https)
	_imageUrlB string
	// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
	_imageType int64
}

// New
