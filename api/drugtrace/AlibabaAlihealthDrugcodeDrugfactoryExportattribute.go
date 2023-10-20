package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportattribute 导出所有项目的药物属性和药品信息
// alibaba.alihealth.drugcode.drugfactory.exportattribute
//
// 导出所有项目的药物属性和药品信息
func AlibabaAlihealthDrugcodeDrugfactoryExportattribute(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
