package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest
非药单码替换 API请求
alibaba.alihealth.tracecodeseller.code.single.codereplace

提供非药追溯码单码替换功能 */
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId string
	// 新码
	_newCode string
	// 老码
	_oldCode string
}

// New
