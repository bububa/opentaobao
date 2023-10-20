package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportproject 导出项目和药品目录
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
func AlibabaAlihealthDrugcodeDrugfactoryExportproject(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
