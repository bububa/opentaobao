package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest
导出所有项目的药物属性和药品信息 API请求
alibaba.alihealth.drugcode.drugfactory.exportattribute

导出所有项目的药物属性和药品信息 */
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// New
