package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseCommonEnumdataAPIRequest
获取通用枚举值接口 API请求
alibaba.legal.case.common.enumdata

获取通用枚举值接口 */
type AlibabaLegalCaseCommonEnumdataAPIRequest struct {
	model.Params
	// bu
	_key string
	// 语言
	_lang string
}

// New
