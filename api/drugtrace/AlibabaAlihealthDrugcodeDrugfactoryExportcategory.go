package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportcategory 导出临床药品目录
// alibaba.alihealth.drugcode.drugfactory.exportcategory
//
// 导出临床药品目录
func AlibabaAlihealthDrugcodeDrugfactoryExportcategory(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
