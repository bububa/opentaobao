package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取盲底文件处理结果 APIRequest
alibaba.alihealth.drugcode.drugfactory.getblindresult

获取盲底文件处理结果
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest struct {
    model.Params

    // 企业id
    refEntId   string 

    // 盲底文件名称
    blindFileName   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.getblindresult"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) SetBlindFileName(blindFileName string) error {
    r.blindFileName = blindFileName
    r.Set("blind_file_name", blindFileName)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetBlindFileName() string {
    return r.blindFileName
}

