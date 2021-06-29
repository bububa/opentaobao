package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gsk新增药店 APIRequest
alibaba.alihealth.alidoc.drug.store.add

GSK上传药店信息
*/
type AlibabaAlihealthAlidocDrugStoreAddRequest struct {
    model.Params

    // 新增药店
    drugStoreAddTopRequest   *DrugStoreAddTopRequest 

}

func NewAlibabaAlihealthAlidocDrugStoreAddRequest() *AlibabaAlihealthAlidocDrugStoreAddRequest{
    return &AlibabaAlihealthAlidocDrugStoreAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alidoc.drug.store.add"
}

func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlidocDrugStoreAddRequest) SetDrugStoreAddTopRequest(drugStoreAddTopRequest *DrugStoreAddTopRequest) error {
    r.drugStoreAddTopRequest = drugStoreAddTopRequest
    r.Set("drug_store_add_top_request", drugStoreAddTopRequest)
    return nil
}

func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetDrugStoreAddTopRequest() *DrugStoreAddTopRequest {
    return r.drugStoreAddTopRequest
}

