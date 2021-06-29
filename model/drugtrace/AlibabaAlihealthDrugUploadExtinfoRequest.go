package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中药饮片及器械对接 APIRequest
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

func NewAlibabaAlihealthDrugUploadExtinfoRequest() *AlibabaAlihealthDrugUploadExtinfoRequest{
    return &AlibabaAlihealthDrugUploadExtinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.upload.extinfo"
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetBatchNo() string {
    return r.batchNo
}

func (r *AlibabaAlihealthDrugUploadExtinfoRequest) SetExtInfoDto(extInfoDto *ExtInfoDto) error {
    r.extInfoDto = extInfoDto
    r.Set("ext_info_dto", extInfoDto)
    return nil
}

func (r AlibabaAlihealthDrugUploadExtinfoRequest) GetExtInfoDto() *ExtInfoDto {
    return r.extInfoDto
}

