package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalBindAuditQueryAPIRequest
ISV查询绑定审核状态 API请求
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态 */
type AlibabaAlihealthDentalBindAuditQueryAPIRequest struct {
	model.Params
	// 绑定ID列表
	_bindIds []int64
}

// New
