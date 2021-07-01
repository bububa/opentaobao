package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrBillcheckAPIRequest
疫苗追溯验证 API请求
alibaba.alihealth.drug.kyt.dr.billcheck

各级疾控在入库完成后，需要做追溯信息验证 */
type AlibabaAlihealthDrugKytDrBillcheckAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 单据编号
	_billCode string
	// 单据类型
	_billType string
	// 单据企业refEntId
	_owerRefEntId string
}

// New
