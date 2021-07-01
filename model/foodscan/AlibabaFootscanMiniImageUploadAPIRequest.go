package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFootscanMiniImageUploadAPIRequest
商家端图片上传 API请求
alibaba.footscan.mini.image.upload

提供图片上传功能，同时进行图片的检测 */
type AlibabaFootscanMiniImageUploadAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *CheckParam
}

// New
