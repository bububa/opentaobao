package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询码是否激活 API请求
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活
*/
type AlibabaAlihealthDrugKytQuerycodeactiveRequest struct {
    model.Params
    // 企业
    _refEntId   string
    // 码
    _codes   []string
}

// 初始化AlibabaAlihealthDrugKytQuerycodeactiveRequest对象
func NewAlibabaAlihealthDrugKytQuerycodeactiveRequest() *AlibabaAlihealthDrugKytQuerycodeactiveRequest{
    return &AlibabaAlihealthDrugKytQuerycodeactiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querycodeactive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业
func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetRefEntId() string {
    return r._refEntId
}
// Codes Setter
// 码
func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetCodes(_codes []string) error {
    r._codes = _codes
    r.Set("codes", _codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetCodes() []string {
    return r._codes
}
