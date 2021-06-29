package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步配置信息 API请求
alibaba.htorder.hotel.sync.config

同步配置信息
*/
type AlibabaHtorderHotelSyncConfigRequest struct {
    model.Params
    // 配置信息
    dataEntity   *HotelMessageConfigDto
}

// 初始化AlibabaHtorderHotelSyncConfigRequest对象
func NewAlibabaHtorderHotelSyncConfigRequest() *AlibabaHtorderHotelSyncConfigRequest{
    return &AlibabaHtorderHotelSyncConfigRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncConfigRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.config"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataEntity Setter
// 配置信息
func (r *AlibabaHtorderHotelSyncConfigRequest) SetDataEntity(dataEntity *HotelMessageConfigDto) error {
    r.dataEntity = dataEntity
    r.Set("data_entity", dataEntity)
    return nil
}

// DataEntity Getter
func (r AlibabaHtorderHotelSyncConfigRequest) GetDataEntity() *HotelMessageConfigDto {
    return r.dataEntity
}
