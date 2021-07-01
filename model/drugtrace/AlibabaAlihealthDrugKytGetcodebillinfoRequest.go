package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取基本和单据信息 API请求
alibaba.alihealth.drug.kyt.getcodebillinfo

根据码信息获取基本信息和单据信息
*/
type AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 码
    _code   string
}

// 初始化AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytGetcodebillinfoRequest() *AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest{
    return &AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getcodebillinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// Code Setter
// 码
func (r *AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest) GetCode() string {
    return r._code
}
