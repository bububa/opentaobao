package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新药店 APIRequest
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口
*/
type AlibabaAlihealthAlidocDrugStoreUpdateRequest struct {
    model.Params

    // 更新对象
    drugStoreUpdateTopRequest   *DrugStoreUpdateTopRequest 

}

func NewAlibabaAlihealthAlidocDrugStoreUpdateRequest() *AlibabaAlihealthAlidocDrugStoreUpdateRequest{
    return &AlibabaAlihealthAlidocDrugStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alidoc.drug.store.update"
}

func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlidocDrugStoreUpdateRequest) SetDrugStoreUpdateTopRequest(drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest) error {
    r.drugStoreUpdateTopRequest = drugStoreUpdateTopRequest
    r.Set("drug_store_update_top_request", drugStoreUpdateTopRequest)
    return nil
}

func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
    return r.drugStoreUpdateTopRequest
}

