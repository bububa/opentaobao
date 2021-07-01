package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest
根据单据编号查询单据明细 API请求
alibaba.alihealth.drug.kyt.query.druginfo.from.billcode

根据单据编号查询单据明细 */
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest struct {
	model.Params
	// 单据号
	_billCode string
	// 企业id
	_refEntId string
}

// New
