package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取盲底文件处理结果 API请求
alibaba.alihealth.drugcode.drugfactory.getblindresult

获取盲底文件处理结果
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest struct {
    model.Params
    // 企业id
    _refEntId   string
    // 盲底文件名称
    _blindFileName   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.getblindresult"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetRefEntId() string {
    return r._refEntId
}
// BlindFileName Setter
// 盲底文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) SetBlindFileName(_blindFileName string) error {
    r._blindFileName = _blindFileName
    r.Set("blind_file_name", _blindFileName)
    return nil
}

// BlindFileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest) GetBlindFileName() string {
    return r._blindFileName
}