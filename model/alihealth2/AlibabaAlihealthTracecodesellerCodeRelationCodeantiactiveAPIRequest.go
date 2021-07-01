package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest
解除码的关联关系 API请求
alibaba.alihealth.tracecodeseller.code.relation.codeantiactive

解除码的关联关系 */
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest struct {
	model.Params
	// 顶层码
	_topCode string
	// 淘宝名
	_tbUserId string
	// 企业id
	_entInfoId int64
}

// New
