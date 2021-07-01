package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingImagetranslateAPIRequest
图片机器翻译 API请求
alibaba.seaking.imagetranslate

图片机器翻译 */
type AlibabaSeakingImagetranslateAPIRequest struct {
	model.Params
	// erp用户id
	_identifier string
	// 目标语种
	_targetLang string
	// 源语种
	_sourceLang string
	// 调用来源(erp名称)
	_identifierType string
	// 原图url
	_url string
	// 扩展信息
	_extra *Extra
}

// New
