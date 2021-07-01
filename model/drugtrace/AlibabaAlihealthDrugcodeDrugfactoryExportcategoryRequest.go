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
type AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportcategoryRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportcategory"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportcategoryAPIRequest) GetRefEntId() string {
    return r._refEntId
}
