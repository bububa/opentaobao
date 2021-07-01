package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePaperCommonReturnAPIRequest
纸票通用回传接口 API请求
alibaba.einvoice.paper.common.return

纸票通用回传接口（打印回传、注册回传等），只返回成功or失败 */
type AlibabaEinvoicePaperCommonReturnAPIRequest struct {
	model.Params
	// 请求索引
	_reqIndex string
	// 回传结果
	_success bool
	// 错误码，success=false时必填
	_bizErrorCode string
	// 错误信息，success=false时必填
	_bizErrorMsg string
	// 扩展信息
	_extProps string
}

// New
