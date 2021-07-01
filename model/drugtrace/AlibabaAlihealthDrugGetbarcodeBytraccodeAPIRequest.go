package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest
根据追溯码获取69码 API请求
alibaba.alihealth.drug.getbarcode.bytraccode

根据追溯码获取69码 */
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
}

// New
