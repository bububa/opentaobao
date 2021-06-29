package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步配置信息 APIRequest
alibaba.htorder.hotel.sync.config

同步配置信息
*/
type AlibabaHtorderHotelSyncConfigRequest struct {
    model.Params

    // 配置信息
    dataEntity   *HotelMessageConfigDto 

}

func NewAlibabaHtorderHotelSyncConfigRequest() *AlibabaHtorderHotelSyncConfigRequest{
    return &AlibabaHtorderHotelSyncConfigRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHtorderHotelSyncConfigRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.config"
}

func (r AlibabaHtorderHotelSyncConfigRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHtorderHotelSyncConfigRequest) SetDataEntity(dataEntity *HotelMessageConfigDto) error {
    r.dataEntity = dataEntity
    r.Set("data_entity", dataEntity)
    return nil
}

func (r AlibabaHtorderHotelSyncConfigRequest) GetDataEntity() *HotelMessageConfigDto {
    return r.dataEntity
}

