package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMarketingFaceSkindetectAPIRequest
肌肤检测 API请求
tmall.marketing.face.skindetect

提供人脸肌肤属性报告 */
type TmallMarketingFaceSkindetectAPIRequest struct {
	model.Params
	// 图片的base64（必须以base64,开头）
	_image string
	// isv标识
	_source string
	// 前置摄像头1，后置摄像头0
	_frontCamera string
	// 混淆nick
	_mixnick string
}

// New
