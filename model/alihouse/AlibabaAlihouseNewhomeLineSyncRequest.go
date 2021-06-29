package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
环线数据同步 APIRequest
alibaba.alihouse.newhome.line.sync

环线数据同步
*/
type AlibabaAlihouseNewhomeLineSyncRequest struct {
    model.Params

    // 环线入参
    baseLoopLineDto   *BaseLoopLineDto 

}

func NewAlibabaAlihouseNewhomeLineSyncRequest() *AlibabaAlihouseNewhomeLineSyncRequest{
    return &AlibabaAlihouseNewhomeLineSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeLineSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.line.sync"
}

func (r AlibabaAlihouseNewhomeLineSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeLineSyncRequest) SetBaseLoopLineDto(baseLoopLineDto *BaseLoopLineDto) error {
    r.baseLoopLineDto = baseLoopLineDto
    r.Set("base_loop_line_dto", baseLoopLineDto)
    return nil
}

func (r AlibabaAlihouseNewhomeLineSyncRequest) GetBaseLoopLineDto() *BaseLoopLineDto {
    return r.baseLoopLineDto
}

