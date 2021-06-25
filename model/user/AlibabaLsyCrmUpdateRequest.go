package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
跟进客资状态接口 APIRequest
alibaba.lsy.crm.update

同步客资状态接口
*/
type AlibabaLsyCrmUpdateRequest struct {
    model.Params

    // 更新客资状态对象
    nrtUpdateRecordStatusDto   *NrtUpdateRecordStatusDto 

}

func NewAlibabaLsyCrmUpdateRequest() *AlibabaLsyCrmUpdateRequest{
    return &AlibabaLsyCrmUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.update"
}

func (r AlibabaLsyCrmUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmUpdateRequest) SetNrtUpdateRecordStatusDto(nrtUpdateRecordStatusDto *NrtUpdateRecordStatusDto) error {
    r.nrtUpdateRecordStatusDto = nrtUpdateRecordStatusDto
    r.Set("nrt_update_record_status_dto", nrtUpdateRecordStatusDto)
    return nil
}

func (r AlibabaLsyCrmUpdateRequest) GetNrtUpdateRecordStatusDto() *NrtUpdateRecordStatusDto {
    return r.nrtUpdateRecordStatusDto
}

