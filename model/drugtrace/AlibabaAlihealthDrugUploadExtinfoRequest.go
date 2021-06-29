package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中药饮片及器械对接 API请求
alibaba.alihealth.drug.upload.extinfo

中药饮片及器械对接
*/
type AlibabaAlihealthDrugUploadExtinfoRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 药品ID
    _drugId   string
    // 批次
    _batchNo   string
    // 扩展信息
    _extInfoDto   *ExtInfoDto
}

// 初始化AlibabaAlihealthDrugUploadExtinfoRequest对象
func NewAlibabaAlihealthDrugUploadExtinfoRequest() *AlibabaAlihealthDrugUploadExtinfoRequest{
    return &AlibabaAlihealthDrugUploadExtinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.upload.extinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetRefEntId() string {
    return r._refEntId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetDrugId() string {
    return r._drugId
}
// BatchNo Setter
// 批次
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetBatchNo(_batchNo string) error {
    r._batchNo = _batchNo
    r.Set("batch_no", _batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetBatchNo() string {
    return r._batchNo
}
// ExtInfoDto Setter
// 扩展信息
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetExtInfoDto(_extInfoDto *ExtInfoDto) error {
    r._extInfoDto = _extInfoDto
    r.Set("ext_info_dto", _extInfoDto)
    return nil
}

// ExtInfoDto Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetExtInfoDto() *ExtInfoDto {
    return r._extInfoDto
}
