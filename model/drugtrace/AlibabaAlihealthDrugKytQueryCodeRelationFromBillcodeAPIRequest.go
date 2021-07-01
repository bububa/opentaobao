package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest
根据单据号码查询码单据详情和码信息 API请求
alibaba.alihealth.drug.kyt.query.code.relation.from.billcode

根据单据号码查询码单据详情和码信息 */
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 企业refEntId
	_refEntId string
}

// New
