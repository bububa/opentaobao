package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDestbillCheckAPIRequest
直调审批 API请求
alibaba.alihealth.drug.kyt.destbill.check

为药企提供直调单据审批操作 */
type AlibabaAlihealthDrugKytDestbillCheckAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 单据号
	_billCode string
	// 审核状态，'Y'审批通过 'N' 审批不通过
	_checkType string
}

// New
