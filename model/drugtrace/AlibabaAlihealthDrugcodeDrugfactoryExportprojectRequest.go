package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出项目和药品目录 API请求
alibaba.alihealth.drugcode.drugfactory.exportproject

导出临床项目及其药品目录
*/
type AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest struct {
    model.Params
    // 企业id
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.exportproject"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportprojectRequest) GetRefEntId() string {
    return r._refEntId
}
