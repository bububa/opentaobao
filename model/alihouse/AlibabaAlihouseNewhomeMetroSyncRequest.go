package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地铁数据同步 API请求
alibaba.alihouse.newhome.metro.sync

地铁数据同步
*/
type AlibabaAlihouseNewhomeMetroSyncAPIRequest struct {
    model.Params
    // 地铁入参数据
    _baseMetroLineDto   *BaseMetroLineDTO
}

// 初始化AlibabaAlihouseNewhomeMetroSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeMetroSyncRequest() *AlibabaAlihouseNewhomeMetroSyncAPIRequest{
    return &AlibabaAlihouseNewhomeMetroSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.metro.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseMetroLineDto Setter
// 地铁入参数据
func (r *AlibabaAlihouseNewhomeMetroSyncAPIRequest) SetBaseMetroLineDto(_baseMetroLineDto *BaseMetroLineDTO) error {
    r._baseMetroLineDto = _baseMetroLineDto
    r.Set("base_metro_line_dto", _baseMetroLineDto)
    return nil
}

// BaseMetroLineDto Getter
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetBaseMetroLineDto() *BaseMetroLineDTO {
    return r._baseMetroLineDto
}
