package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品资质信息 API请求
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。
*/
type AlibabaAlihealthDrugKytGetdruglicenseRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 药品ID
    _drugId   string
}

// 初始化AlibabaAlihealthDrugKytGetdruglicenseRequest对象
func NewAlibabaAlihealthDrugKytGetdruglicenseRequest() *AlibabaAlihealthDrugKytGetdruglicenseRequest{
    return &AlibabaAlihealthDrugKytGetdruglicenseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getdruglicense"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetRefEntId() string {
    return r._refEntId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetDrugId() string {
    return r._drugId
}
