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
    refEntId   string
    // 药品ID
    drugId   string
    // 批次
    batchNo   string
    // 扩展信息
    extInfoDto   *ExtInfoDto
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
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetRefEntId() string {
    return r.refEntId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetDrugId() string {
    return r.drugId
}
// BatchNo Setter
// 批次
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetBatchNo() string {
    return r.batchNo
}
// ExtInfoDto Setter
// 扩展信息
func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetExtInfoDto(extInfoDto *ExtInfoDto) error {
    r.extInfoDto = extInfoDto
    r.Set("ext_info_dto", extInfoDto)
    return nil
}

// ExtInfoDto Getter
func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetExtInfoDto() *ExtInfoDto {
    return r.extInfoDto
}
