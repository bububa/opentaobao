package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest
根据追溯码修改69码 API请求
alibaba.alihealth.drug.updatebarcode.bytraccode

根据追溯码修改69码 */
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
	// 69码
	_barcode string
}

// New
