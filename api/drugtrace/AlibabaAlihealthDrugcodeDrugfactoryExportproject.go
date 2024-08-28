package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportproject 导出项目和药品目录
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
func AlibabaAlihealthDrugcodeDrugfactoryExportproject(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
