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
type AlibabaAlihouseNewhomeLineSyncRequest struct {
    model.Params
    // 环线入参
    _baseLoopLineDto   *BaseLoopLineDto
}

// 初始化AlibabaAlihouseNewhomeLineSyncRequest对象
func NewAlibabaAlihouseNewhomeLineSyncRequest() *AlibabaAlihouseNewhomeLineSyncRequest{
    return &AlibabaAlihouseNewhomeLineSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLineSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.line.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLineSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseLoopLineDto Setter
// 环线入参
func (r *AlibabaAlihouseNewhomeLineSyncRequest) SetBaseLoopLineDto(_baseLoopLineDto *BaseLoopLineDto) error {
    r._baseLoopLineDto = _baseLoopLineDto
    r.Set("base_loop_line_dto", _baseLoopLineDto)
    return nil
}

// BaseLoopLineDto Getter
func (r AlibabaAlihouseNewhomeLineSyncRequest) GetBaseLoopLineDto() *BaseLoopLineDto {
    return r._baseLoopLineDto
}
