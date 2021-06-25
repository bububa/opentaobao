package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建客资 APIRequest
alibaba.lsy.crm.create

欢客调用该接口进行客资创建
*/
type AlibabaLsyCrmCreateRequest struct {
    model.Params

    // 客资记录对象
    nrtRecordDto   *NrtRecordDto 

}

func NewAlibabaLsyCrmCreateRequest() *AlibabaLsyCrmCreateRequest{
    return &AlibabaLsyCrmCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmCreateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.create"
}

func (r AlibabaLsyCrmCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmCreateRequest) SetNrtRecordDto(nrtRecordDto *NrtRecordDto) error {
    r.nrtRecordDto = nrtRecordDto
    r.Set("nrt_record_dto", nrtRecordDto)
    return nil
}

func (r AlibabaLsyCrmCreateRequest) GetNrtRecordDto() *NrtRecordDto {
    return r.nrtRecordDto
}

