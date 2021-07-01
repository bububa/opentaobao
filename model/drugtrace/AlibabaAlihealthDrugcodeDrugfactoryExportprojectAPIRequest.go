package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest
导出项目和药品目录 API请求
alibaba.alihealth.drugcode.drugfactory.exportproject

导出临床项目及其药品目录 */
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// New
