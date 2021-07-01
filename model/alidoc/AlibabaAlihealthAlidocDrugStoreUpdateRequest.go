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
type AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest struct {
    model.Params
    // 更新对象
    _drugStoreUpdateTopRequest   *DrugStoreUpdateTopRequest
}

// 初始化AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest对象
func NewAlibabaAlihealthAlidocDrugStoreUpdateRequest() *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest{
    return &AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alidoc.drug.store.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DrugStoreUpdateTopRequest Setter
// 更新对象
func (r *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) SetDrugStoreUpdateTopRequest(_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest) error {
    r._drugStoreUpdateTopRequest = _drugStoreUpdateTopRequest
    r.Set("drug_store_update_top_request", _drugStoreUpdateTopRequest)
    return nil
}

// DrugStoreUpdateTopRequest Getter
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
    return r._drugStoreUpdateTopRequest
}
