package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingDiagnosistitleAPIRequest
标题诊断 API请求
alibaba.seaking.diagnosistitle

标题诊断 */
type AlibabaSeakingDiagnosistitleAPIRequest struct {
	model.Params
	// 类目id,没有的时候传-1
	_categoryId int64
	// 扩展信息
	_extra *Extra
	// erp用户id
	_identifier string
	// 调用来源(erp名称)
	_identifierType string
	// 语种
	_language string
	// 商品所在平台（ae/icbu）
	_platform string
	// 标题
	_title string
}

// New
