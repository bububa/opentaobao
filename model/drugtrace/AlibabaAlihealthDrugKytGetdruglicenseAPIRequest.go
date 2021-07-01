package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetdruglicenseAPIRequest
获取药品资质信息 API请求
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。 */
type AlibabaAlihealthDrugKytGetdruglicenseAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugId string
}

// New
