package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
导出临床药品目录 
alibaba.alihealth.drugcode.drugfactory.exportcategory

导出临床药品目录
*/
func AlibabaAlihealthDrugcodeDrugfactoryExportcategory(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
