package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingTranslateAPIRequest
MT定制接口 API请求
alibaba.seaking.translate

MT定制接口 */
type AlibabaSeakingTranslateAPIRequest struct {
	model.Params
	// 定制用户id
	_identifier string
	// 目标语种
	_targetLang string
	// 源语种
	_sourceLang string
	// 原文
	_sourceText string
	// 原文格式(text/html)
	_sourceFormat string
	// 定制用户类型
	_identifierType string
	// 原文类型(title: 标题/offer: 详描/message: 消息)
	_fieldType string
	// 扩展信息
	_extra *Extra
}

// New
