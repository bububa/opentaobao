package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出临床药品目录 API请求
alibaba.alihealth.drugcode.drugfactory.exportcategory

导出临床药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest struct {
    model.Params
    // 企业ID
    refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportcategory"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest) GetRefEntId() string {
    return r.refEntId
}
