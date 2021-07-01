package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytCodetobillAPIRequest
通过追溯码查单据 API请求
alibaba.alihealth.drug.kyt.codetobill

通过追溯码查单据 */
type AlibabaAlihealthDrugKytCodetobillAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 追溯码
	_code string
}

// New
