package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv协议获取 APIRequest
alibaba.alihealth.examination.agreement.list

isv协议获取
*/
type AlibabaAlihealthExaminationAgreementListRequest struct {
    model.Params

    // isv传递过来的门店code
    storeCode   string 

}

func NewAlibabaAlihealthExaminationAgreementListRequest() *AlibabaAlihealthExaminationAgreementListRequest{
    return &AlibabaAlihealthExaminationAgreementListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationAgreementListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.agreement.list"
}

func (r AlibabaAlihealthExaminationAgreementListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationAgreementListRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r AlibabaAlihealthExaminationAgreementListRequest) GetStoreCode() string {
    return r.storeCode
}

