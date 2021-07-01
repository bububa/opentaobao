package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商圈数据同步 API请求
alibaba.alihouse.newhome.business.sync

商圈数据同步
*/
type AlibabaAlihouseNewhomeBusinessSyncAPIRequest struct {
    model.Params
    // 入参数据
    _baseBusinessDto   *BaseBusinessDTO
}

// 初始化AlibabaAlihouseNewhomeBusinessSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeBusinessSyncRequest() *AlibabaAlihouseNewhomeBusinessSyncAPIRequest{
    return &AlibabaAlihouseNewhomeBusinessSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.business.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseBusinessDto Setter
// 入参数据
func (r *AlibabaAlihouseNewhomeBusinessSyncAPIRequest) SetBaseBusinessDto(_baseBusinessDto *BaseBusinessDTO) error {
    r._baseBusinessDto = _baseBusinessDto
    r.Set("base_business_dto", _baseBusinessDto)
    return nil
}

// BaseBusinessDto Getter
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetBaseBusinessDto() *BaseBusinessDTO {
    return r._baseBusinessDto
}
