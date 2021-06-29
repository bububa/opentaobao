package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地铁数据同步 APIRequest
alibaba.alihouse.newhome.metro.sync

地铁数据同步
*/
type AlibabaAlihouseNewhomeMetroSyncRequest struct {
    model.Params

    // 地铁入参数据
    baseMetroLineDto   *BaseMetroLineDto 

}

func NewAlibabaAlihouseNewhomeMetroSyncRequest() *AlibabaAlihouseNewhomeMetroSyncRequest{
    return &AlibabaAlihouseNewhomeMetroSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.metro.sync"
}

func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeMetroSyncRequest) SetBaseMetroLineDto(baseMetroLineDto *BaseMetroLineDto) error {
    r.baseMetroLineDto = baseMetroLineDto
    r.Set("base_metro_line_dto", baseMetroLineDto)
    return nil
}

func (r AlibabaAlihouseNewhomeMetroSyncRequest) GetBaseMetroLineDto() *BaseMetroLineDto {
    return r.baseMetroLineDto
}

