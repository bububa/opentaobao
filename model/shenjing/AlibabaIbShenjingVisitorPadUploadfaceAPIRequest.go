package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbShenjingVisitorPadUploadfaceAPIRequest
访客PAD上传人脸 API请求
alibaba.ib.shenjing.visitor.pad.uploadface

访客PAD端上传人脸。 */
type AlibabaIbShenjingVisitorPadUploadfaceAPIRequest struct {
	model.Params
	// 访客ID
	_id string
	// 图片URL
	_image string
}

// New
