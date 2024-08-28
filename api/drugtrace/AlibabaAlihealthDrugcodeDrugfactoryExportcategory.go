package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportcategory 导出临床药品目录
// alibaba.alihealth.drugcode.drugfactory.exportcategory
//
// 导出临床药品目录
func AlibabaAlihealthDrugcodeDrugfactoryExportcategory(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
