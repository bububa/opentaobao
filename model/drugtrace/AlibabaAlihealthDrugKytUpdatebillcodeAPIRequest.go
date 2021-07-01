package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest
零售修改出入库单追溯码 API请求
alibaba.alihealth.drug.kyt.updatebillcode

零售修改出入库单追溯码 */
type AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 操作人ID
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
	// 追溯码
	_codeList []string
}

// New
