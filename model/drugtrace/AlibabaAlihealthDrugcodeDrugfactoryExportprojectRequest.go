package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出项目和药品目录 APIRequest
alibaba.alihealth.drugcode.drugfactory.exportproject

导出临床项目及其药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest struct {
    model.Params

    // 企业id
    refEntId   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportproject"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetRefEntId() string {
    return r.refEntId
}

