package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest
家居属性识别 API请求
aliyun.viapi.goodstech.recognize.furniture.attribute

识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiGoodstechRecognizeFurnitureAttributeAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New
