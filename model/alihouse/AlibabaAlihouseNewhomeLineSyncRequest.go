package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
环线数据同步 API请求
alibaba.alihouse.newhome.line.sync

环线数据同步
*/
type AlibabaAlihouseNewhomeLineSyncAPIRequest struct {
    model.Params
    // 环线入参
    _baseLoopLineDto   *BaseLoopLineDTO
}

// 初始化AlibabaAlihouseNewhomeLineSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeLineSyncRequest() *AlibabaAlihouseNewhomeLineSyncAPIRequest{
    return &AlibabaAlihouseNewhomeLineSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.line.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseLoopLineDto Setter
// 环线入参
func (r *AlibabaAlihouseNewhomeLineSyncAPIRequest) SetBaseLoopLineDto(_baseLoopLineDto *BaseLoopLineDTO) error {
    r._baseLoopLineDto = _baseLoopLineDto
    r.Set("base_loop_line_dto", _baseLoopLineDto)
    return nil
}

// BaseLoopLineDto Getter
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetBaseLoopLineDto() *BaseLoopLineDTO {
    return r._baseLoopLineDto
}
