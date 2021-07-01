package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalStoreAuditQueryAPIRequest
ISV查询门店审核状态 API请求
alibaba.alihealth.dental.store.audit.query

ISV查询门店审核状态 */
type AlibabaAlihealthDentalStoreAuditQueryAPIRequest struct {
	model.Params
	// 审核ID列表
	_storeAuditIds []int64
}

// New
