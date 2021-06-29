package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过追溯码查单据 API请求
alibaba.alihealth.drug.kyt.codetobill

通过追溯码查单据
*/
type AlibabaAlihealthDrugKytCodetobillRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 追溯码
    code   string
}

// 初始化AlibabaAlihealthDrugKytCodetobillRequest对象
func NewAlibabaAlihealthDrugKytCodetobillRequest() *AlibabaAlihealthDrugKytCodetobillRequest{
    return &AlibabaAlihealthDrugKytCodetobillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodetobillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.codetobill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodetobillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytCodetobillRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytCodetobillRequest) GetRefEntId() string {
    return r.refEntId
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytCodetobillRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytCodetobillRequest) GetCode() string {
    return r.code
}
