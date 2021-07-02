package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportproject 导出项目和药品目录
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
func AlibabaAlihealthDrugcodeDrugfactoryExportproject(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
