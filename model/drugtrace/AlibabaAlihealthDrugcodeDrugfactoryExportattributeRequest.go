package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出所有项目的药物属性和药品信息 APIRequest
alibaba.alihealth.drugcode.drugfactory.exportattribute

导出所有项目的药物属性和药品信息
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest struct {
    model.Params

    // 企业id
    refEntId   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportattribute"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetRefEntId() string {
    return r.refEntId
}

