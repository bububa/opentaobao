package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城区数据同步 APIRequest
alibaba.alihouse.newhome.region.sync

城区数据同步
*/
type AlibabaAlihouseNewhomeRegionSyncRequest struct {
    model.Params

    // 城区数据
    baseRegionDto   *BaseRegionDto 

}

func NewAlibabaAlihouseNewhomeRegionSyncRequest() *AlibabaAlihouseNewhomeRegionSyncRequest{
    return &AlibabaAlihouseNewhomeRegionSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeRegionSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.region.sync"
}

func (r AlibabaAlihouseNewhomeRegionSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeRegionSyncRequest) SetBaseRegionDto(baseRegionDto *BaseRegionDto) error {
    r.baseRegionDto = baseRegionDto
    r.Set("base_region_dto", baseRegionDto)
    return nil
}

func (r AlibabaAlihouseNewhomeRegionSyncRequest) GetBaseRegionDto() *BaseRegionDto {
    return r.baseRegionDto
}

