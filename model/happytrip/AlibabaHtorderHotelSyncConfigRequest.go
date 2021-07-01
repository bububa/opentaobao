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
type AlibabaHtorderHotelSyncConfigAPIRequest struct {
    model.Params
    // 配置信息
    _dataEntity   *HotelMessageConfigDTO
}

// 初始化AlibabaHtorderHotelSyncConfigAPIRequest对象
func NewAlibabaHtorderHotelSyncConfigRequest() *AlibabaHtorderHotelSyncConfigAPIRequest{
    return &AlibabaHtorderHotelSyncConfigAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.config"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataEntity Setter
// 配置信息
func (r *AlibabaHtorderHotelSyncConfigAPIRequest) SetDataEntity(_dataEntity *HotelMessageConfigDTO) error {
    r._dataEntity = _dataEntity
    r.Set("data_entity", _dataEntity)
    return nil
}

// DataEntity Getter
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetDataEntity() *HotelMessageConfigDTO {
    return r._dataEntity
}
