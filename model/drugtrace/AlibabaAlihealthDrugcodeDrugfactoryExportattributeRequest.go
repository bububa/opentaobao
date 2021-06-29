package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出所有项目的药物属性和药品信息 API请求
alibaba.alihealth.drugcode.drugfactory.exportattribute

导出所有项目的药物属性和药品信息
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest struct {
    model.Params
    // 企业id
    refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportattribute"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest) GetRefEntId() string {
    return r.refEntId
}
