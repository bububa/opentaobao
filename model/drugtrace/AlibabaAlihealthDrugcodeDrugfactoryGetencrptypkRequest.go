package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取加密公钥 APIRequest
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest struct {
    model.Params

    // 企业Id
    refEntId   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.getencrptypk"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetRefEntId() string {
    return r.refEntId
}

