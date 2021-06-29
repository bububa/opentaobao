package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gsk新增药店 API请求
alibaba.alihealth.alidoc.drug.store.add

GSK上传药店信息
*/
type AlibabaAlihealthAlidocDrugStoreAddRequest struct {
    model.Params
    // 新增药店
    drugStoreAddTopRequest   *DrugStoreAddTopRequest
}

// 初始化AlibabaAlihealthAlidocDrugStoreAddRequest对象
func NewAlibabaAlihealthAlidocDrugStoreAddRequest() *AlibabaAlihealthAlidocDrugStoreAddRequest{
    return &AlibabaAlihealthAlidocDrugStoreAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alidoc.drug.store.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DrugStoreAddTopRequest Setter
// 新增药店
func (r *AlibabaAlihealthAlidocDrugStoreAddRequest) SetDrugStoreAddTopRequest(drugStoreAddTopRequest *DrugStoreAddTopRequest) error {
    r.drugStoreAddTopRequest = drugStoreAddTopRequest
    r.Set("drug_store_add_top_request", drugStoreAddTopRequest)
    return nil
}

// DrugStoreAddTopRequest Getter
func (r AlibabaAlihealthAlidocDrugStoreAddRequest) GetDrugStoreAddTopRequest() *DrugStoreAddTopRequest {
    return r.drugStoreAddTopRequest
}
