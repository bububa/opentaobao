package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest
导出临床药品目录 API请求
alibaba.alihealth.drugcode.drugfactory.exportcategory

导出临床药品目录 */
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
}

// New
