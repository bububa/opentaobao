package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv协议获取 API请求
alibaba.alihealth.examination.agreement.list

isv协议获取
*/
type AlibabaAlihealthExaminationAgreementListRequest struct {
    model.Params
    // isv传递过来的门店code
    _storeCode   string
}

// 初始化AlibabaAlihealthExaminationAgreementListRequest对象
func NewAlibabaAlihealthExaminationAgreementListRequest() *AlibabaAlihealthExaminationAgreementListRequest{
    return &AlibabaAlihealthExaminationAgreementListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationAgreementListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.agreement.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationAgreementListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreCode Setter
// isv传递过来的门店code
func (r *AlibabaAlihealthExaminationAgreementListRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaAlihealthExaminationAgreementListRequest) GetStoreCode() string {
    return r._storeCode
}
