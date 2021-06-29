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
type AlibabaAlihouseNewhomeBusinessSyncRequest struct {
    model.Params
    // 入参数据
    baseBusinessDto   *BaseBusinessDto
}

// 初始化AlibabaAlihouseNewhomeBusinessSyncRequest对象
func NewAlibabaAlihouseNewhomeBusinessSyncRequest() *AlibabaAlihouseNewhomeBusinessSyncRequest{
    return &AlibabaAlihouseNewhomeBusinessSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.business.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BaseBusinessDto Setter
// 入参数据
func (r *AlibabaAlihouseNewhomeBusinessSyncRequest) SetBaseBusinessDto(baseBusinessDto *BaseBusinessDto) error {
    r.baseBusinessDto = baseBusinessDto
    r.Set("base_business_dto", baseBusinessDto)
    return nil
}

// BaseBusinessDto Getter
func (r AlibabaAlihouseNewhomeBusinessSyncRequest) GetBaseBusinessDto() *BaseBusinessDto {
    return r.baseBusinessDto
}
