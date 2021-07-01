package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取上游温度信息（疫苗） API请求
alibaba.alihealth.drug.kyt.dr.getupteminfo

根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
*/
type AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 企业ID
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetupteminfoRequest() *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest{
    return &AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getupteminfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetRefEntId() string {
    return r._refEntId
}
