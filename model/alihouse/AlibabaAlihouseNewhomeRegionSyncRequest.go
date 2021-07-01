package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
城区数据同步 API请求
alibaba.alihouse.newhome.region.sync

城区数据同步
*/
type AlibabaAlihouseNewhomeRegionSyncAPIRequest struct {
    model.Params
    // 城区数据
    _baseRegionDto   *BaseRegionDTO
}

// 初始化AlibabaAlihouseNewhomeRegionSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeRegionSyncRequest() *AlibabaAlihouseNewhomeRegionSyncAPIRequest{
    return &AlibabaAlihouseNewhomeRegionSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.region.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseRegionDto Setter
// 城区数据
func (r *AlibabaAlihouseNewhomeRegionSyncAPIRequest) SetBaseRegionDto(_baseRegionDto *BaseRegionDTO) error {
    r._baseRegionDto = _baseRegionDto
    r.Set("base_region_dto", _baseRegionDto)
    return nil
}

// BaseRegionDto Getter
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetBaseRegionDto() *BaseRegionDTO {
    return r._baseRegionDto
}
