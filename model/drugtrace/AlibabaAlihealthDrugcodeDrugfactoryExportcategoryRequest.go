package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出临床药品目录 APIRequest
alibaba.alihealth.drugcode.drugfactory.exportcategory

导出临床药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportcategory"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetRefEntId() string {
    return r.refEntId
}

