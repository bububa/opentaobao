package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商圈数据同步 APIRequest
alibaba.alihouse.newhome.business.sync

商圈数据同步
*/
type AlibabaAlihouseNewhomeBusinessSyncRequest struct {
    model.Params

    // 入参数据
    baseBusinessDto   *BaseBusinessDto 

}

func NewAlibabaAlihouseNewhomeBusinessSyncRequest() *AlibabaAlihouseNewhomeBusinessSyncRequest{
    return &AlibabaAlihouseNewhomeBusinessSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.business.sync"
}

func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeBusinessSyncRequest) SetBaseBusinessDto(baseBusinessDto *BaseBusinessDto) error {
    r.baseBusinessDto = baseBusinessDto
    r.Set("base_business_dto", baseBusinessDto)
    return nil
}

func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetBaseBusinessDto() *BaseBusinessDto {
    return r.baseBusinessDto
}

