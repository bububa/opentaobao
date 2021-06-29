package alidoc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新药店 API请求
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口
*/
type AlibabaAlihealthAlidocDrugStoreUpdateRequest struct {
    model.Params
    // 更新对象
    _drugStoreUpdateTopRequest   *DrugStoreUpdateTopRequest
}

// 初始化AlibabaAlihealthAlidocDrugStoreUpdateRequest对象
func NewAlibabaAlihealthAlidocDrugStoreUpdateRequest() *AlibabaAlihealthAlidocDrugStoreUpdateRequest{
    return &AlibabaAlihealthAlidocDrugStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alidoc.drug.store.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DrugStoreUpdateTopRequest Setter
// 更新对象
func (r *AlibabaAlihealthAlidocDrugStoreUpdateRequest) SetDrugStoreUpdateTopRequest(_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest) error {
    r._drugStoreUpdateTopRequest = _drugStoreUpdateTopRequest
    r.Set("drug_store_update_top_request", _drugStoreUpdateTopRequest)
    return nil
}

// DrugStoreUpdateTopRequest Getter
func (r AlibabaAlihealthAlidocDrugStoreUpdateRequest) GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
    return r._drugStoreUpdateTopRequest
}
