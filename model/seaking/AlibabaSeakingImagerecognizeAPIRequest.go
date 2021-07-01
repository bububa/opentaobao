package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingImagerecognizeAPIRequest
图片语种识别 API请求
alibaba.seaking.imagerecognize

图片语种识别 */
type AlibabaSeakingImagerecognizeAPIRequest struct {
	model.Params
	// 扩展信息
	_extra *Extra
	// erp用户id
	_identifier string
	// 调用来源(erp名称)
	_identifierType string
	// 图片url
	_url string
}

// New
