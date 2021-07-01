package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetentlicenseAPIRequest
获取企业资质 API请求
alibaba.alihealth.drug.kyt.getentlicense

获取企业的资质信息。 */
type AlibabaAlihealthDrugKytGetentlicenseAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
}

// New
