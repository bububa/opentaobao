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
    refEntId   string
    // 药品ID
    drugId   string
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
func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetRefEntId() string {
    return r.refEntId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetDrugId() string {
    return r.drugId
}
